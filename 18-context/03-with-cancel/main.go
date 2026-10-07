package main

import (
	"context"
	"fmt"
)

// worker handles jobs until the caller cancels the context. ctx is always the
// first parameter and is named ctx by convention. The arrow on a channel type
// marks its direction: <-chan is receive-only, chan<- is send-only.
func worker(ctx context.Context, jobs <-chan string, report chan<- string) {
	for {
		// select waits on whichever channel is ready first. Here the worker
		// either gets the next job or notices that the context was cancelled.
		select {
		case <-ctx.Done(): // Closed when the caller calls cancel().
			report <- "worker: cancelled, stopping"
			return
		case job := <-jobs: // <-jobs receives one job value.
			report <- "worker: handled " + job
		}
	}
}

func main() {
	// WithCancel returns a new context plus a cancel function. Calling cancel
	// is how the caller signals "stop".
	ctx, cancel := context.WithCancel(context.Background())

	jobs := make(chan string)   // Unbuffered: a send waits for the worker.
	report := make(chan string) // The worker reports each outcome back.
	go worker(ctx, jobs, report)

	jobs <- "restock tea" // Hand the worker one job...
	fmt.Println(<-report) // ...and wait for its report.

	cancel()              // Now tell the worker to stop.
	fmt.Println(<-report) // No more jobs are sent, so select picks ctx.Done().
}
