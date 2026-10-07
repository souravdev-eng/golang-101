package main

import (
	"context"
	"fmt"
	"time"
)

// fetch pretends to call a slow service. It reports its result on a buffered
// channel so it never blocks, even if nobody is waiting anymore. chan<- string
// is a send-only channel: fetch may send on it but not receive.
func fetch(ctx context.Context, result chan<- string) {
	// time.AfterFunc runs the function after the delay. One full second is far
	// longer than the timeout below, so the timeout always wins.
	time.AfterFunc(1*time.Second, func() { result <- "inventory data" })
}

func main() {
	// WithTimeout cancels the context automatically after the duration. It also
	// returns a cancel function.
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel() // Always release the context's resources, even on timeout.

	result := make(chan string, 1) // Buffered (size 1) so fetch never blocks.
	go fetch(ctx, result)

	select {
	case data := <-result: // Would win only if the work beat the timeout.
		fmt.Println("got:", data)
	case <-ctx.Done(): // Fires after 20ms because the work needs a full second.
		fmt.Println("gave up:", ctx.Err())
	}
}
