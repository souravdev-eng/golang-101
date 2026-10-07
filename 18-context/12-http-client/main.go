package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

// On the client side, you attach a context to an outbound request with
// http.NewRequestWithContext. If the context is cancelled or its deadline
// passes before the reply arrives, the in-flight call is aborted — you are not
// stuck waiting on a slow or hung service.

// slowHandler stands in for a sluggish downstream service: it takes a full
// second to answer. It also watches its own request context so that when the
// client disconnects, it returns at once instead of blocking srv.Close().
func slowHandler(w http.ResponseWriter, r *http.Request) {
	select {
	case <-time.After(1 * time.Second):
		fmt.Fprintln(w, "inventory data")
	case <-r.Context().Done():
	}
}

func main() {
	// A local test server, so the example needs no real network.
	srv := httptest.NewServer(http.HandlerFunc(slowHandler))
	defer srv.Close()

	// Allow only 20ms for a call the server answers in a full second, so the
	// deadline always fires first and the request is aborted.
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	// The context rides along with the request; the http.Client enforces it.
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, srv.URL, nil)
	_, err := http.DefaultClient.Do(req)

	// The returned error's text embeds the server's random port, so report the
	// cause from ctx.Err() instead, which is a stable value.
	if err != nil {
		fmt.Println("client: aborted the call:", ctx.Err())
	}
}
