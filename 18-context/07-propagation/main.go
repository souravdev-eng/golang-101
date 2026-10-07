package main

import (
	"context"
	"fmt"
)

// One context is created at the top and passed straight down the call chain:
// stepA -> stepB -> stepC. Each function takes ctx as its first parameter and
// hands the SAME ctx to the next. Cancelling once at the top is felt all the way
// down at the bottom, because they all watch the same Done() channel.

// stepC is the bottom of the chain. It does the actual waiting. Before starting
// its "work" it checks whether the context is already cancelled.
func stepC(ctx context.Context) {
	select {
	case <-ctx.Done(): // Closed when the top-level cancel() was called.
		fmt.Println("stepC: ctx cancelled, not starting work")
	default: // default runs when no other case is ready.
		fmt.Println("stepC: doing work")
	}
}

// stepB sits in the middle. It does no context work of its own; it just passes
// ctx along. This is the common case: most functions only forward the context.
func stepB(ctx context.Context) {
	fmt.Println("stepB: passing ctx down to stepC")
	stepC(ctx)
}

// stepA is the top of the chain inside the call tree.
func stepA(ctx context.Context) {
	fmt.Println("stepA: passing ctx down to stepB")
	stepB(ctx)
}

func main() {
	// Create one context for the whole chain and cancel it immediately, so by
	// the time stepC checks, the signal has already propagated all the way down.
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel at the top...

	stepA(ctx) // ...and the bottom of the chain sees it.
}
