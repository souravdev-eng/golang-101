package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// A deadline is a specific moment in time, not a length of time. WithTimeout
	// says "20ms from now"; WithDeadline says "at exactly this instant". Here we
	// build that instant by adding 20ms to the current time.
	deadline := time.Now().Add(20 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel() // Release resources even though the deadline will fire first.

	result := make(chan string, 1) // Buffered so the goroutine never blocks.
	go func() {
		time.Sleep(1 * time.Second) // The work takes far longer than we allow.
		result <- "report ready"
	}()

	select {
	case data := <-result: // Would win only if the work beat the deadline.
		fmt.Println("got:", data)
	case <-ctx.Done(): // Fires the moment `deadline` passes.
		fmt.Println("missed the deadline:", ctx.Err())
	}
}
