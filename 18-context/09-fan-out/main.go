package main

import (
	"context"
	"fmt"
	"sort"
	"sync"
)

// Fan-out: one context is shared by several workers at once. A single cancel()
// stops all of them, because they all watch the same ctx.Done(). Here one
// worker hits an error and triggers the cancel; the others notice and stop.
//
// Standard library only: sync.WaitGroup to wait for the workers, and a shared
// context to signal them. (The golang.org/x/sync/errgroup package bundles this
// exact pattern, but we stay in the standard library to show the moving parts.)

// worker loops until either it is told to stop (ctx.Done()) or it decides to
// fail. The worker whose id matches failAt reports an error and calls cancel,
// which stops every other worker sharing this ctx. Results are sent back on the
// out channel instead of printed, so main can order them deterministically.
func worker(ctx context.Context, id, failAt int, cancel context.CancelFunc, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	if id == failAt {
		out <- fmt.Sprintf("worker %d: failed, cancelling the group", id)
		cancel() // Signal every worker sharing ctx to stop.
		return
	}

	// Not the failing worker: wait for the shared cancel to arrive.
	<-ctx.Done()
	out <- fmt.Sprintf("worker %d: stopped (%v)", id, ctx.Err())
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Release the context even if no worker fails.

	const workers = 3
	const failAt = 2 // Worker 2 will fail and cancel the rest.

	out := make(chan string, workers) // Buffered so no worker blocks on send.
	var wg sync.WaitGroup
	for id := 1; id <= workers; id++ {
		wg.Add(1)
		go worker(ctx, id, failAt, cancel, out, &wg)
	}

	wg.Wait()  // Wait for all workers to finish.
	close(out) // Safe to close now: every sender has returned.

	// Concurrency makes the send order unpredictable, so collect and sort the
	// lines before printing. Sorting gives the same output on every run.
	var lines []string
	for line := range out {
		lines = append(lines, line)
	}
	sort.Strings(lines)
	for _, line := range lines {
		fmt.Println(line)
	}
}
