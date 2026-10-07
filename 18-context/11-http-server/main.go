package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

// On the server side, every *http.Request carries a context you read with
// r.Context(). It is cancelled when the client goes away — the browser closes
// the tab, the connection drops, or the client's own timeout fires. Watching it
// lets a handler abandon expensive work nobody is waiting for anymore.

func main() {
	// outcome carries the handler's result back to main so main can print both
	// lines in order. Without it the handler (running in the server's goroutine)
	// and main would race to stdout and the output would not be deterministic.
	// Keeping it a local channel the handler closure captures mirrors the rest
	// of the section, where channels are created in main and passed to the work.
	outcome := make(chan string, 1)

	// racingHandler pretends to assemble an inventory report that takes a full
	// second. It races that slow work against r.Context().Done(), which fires
	// when the client disconnects, and reports which one won on outcome.
	racingHandler := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		select {
		case <-time.After(1 * time.Second): // The slow work finishing first.
			fmt.Fprintln(w, "inventory ready")
			outcome <- "handler: finished the slow work"
		case <-ctx.Done(): // The client gave up; stop wasting effort on it.
			outcome <- fmt.Sprintf("handler: client gone, abandoning slow work: %v", ctx.Err())
		}
	}

	// httptest.NewServer spins up a real HTTP server on a random local port, so
	// the example needs no network and cleans up with srv.Close().
	srv := httptest.NewServer(http.HandlerFunc(racingHandler))
	defer srv.Close()

	// The client allows only 20ms, far less than the handler's one-second work,
	// so the client's context deadline always fires first and it disconnects.
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, srv.URL, nil)
	_, err := http.DefaultClient.Do(req)

	// The raw error text includes the random port, so report a fixed message
	// after confirming the cause, keeping the output deterministic.
	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("client: timed out waiting for the server")
	}

	// Wait for the handler to notice the disconnect and report, then print it.
	fmt.Println(<-outcome)
}
