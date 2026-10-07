package main

import (
	"fmt"
	"sync" // WaitGroup lets main wait for a goroutine to finish.
)

// countUp does a fixed amount of work. Notice what it is missing: there is no
// input that means "stop early". Once it starts, it runs every step, even if
// the caller stops caring after the first one. That missing stop signal is the
// problem the context package solves.
func countUp(steps int, workers *sync.WaitGroup) {
	defer workers.Done()                   // Mark the work done when this returns.
	for step := 1; step <= steps; step++ { // := declares-and-assigns the counter.
		fmt.Println("worker: step", step)
	}
}

func main() {
	var workers sync.WaitGroup
	workers.Add(1)          // Register one unit of work before starting it.
	go countUp(5, &workers) // go runs countUp in a separate goroutine.

	// We have no way to say "stop at step 1". The only option is to wait for
	// every step to finish. Wait() blocks until the registered work is done,
	// which keeps the printed order predictable.
	workers.Wait()
	fmt.Println("main: the worker ran all 5 steps; I had no way to stop it")
}
