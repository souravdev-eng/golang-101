package main

import (
	"fmt"
	"sync" // WaitGroup tracks work that must finish before continuing.
)

func main() {
	var workers sync.WaitGroup
	workers.Add(1) // Register the work before starting the goroutine.
	go func() {    // go starts this function in another goroutine.
		defer workers.Done() // Mark completion when the function returns.
		fmt.Println("worker: packed lunch")
	}()
	workers.Wait() // Block until the registered work has finished.
	fmt.Println("main: ready to leave")
	// No sleep is needed; Wait establishes the completion order.
}
