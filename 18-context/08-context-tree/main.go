package main

import (
	"context"
	"fmt"
)

// Contexts form a tree. A child is DERIVED from a parent with WithCancel (or
// WithTimeout, WithDeadline). Cancellation flows DOWN the tree only:
//   - cancel the parent  -> the child is cancelled too.
//   - cancel the child   -> the parent stays alive.
// This example shows both directions with two separate parent/child pairs.

// state returns "done" or "alive" by peeking at whether a context is cancelled.
// It never blocks: the select has a default branch for the "not cancelled" case.
func state(ctx context.Context) string {
	select {
	case <-ctx.Done():
		return "done"
	default:
		return "alive"
	}
}

func main() {
	// Pair 1: cancelling the parent cancels the child.
	parent1, cancelParent1 := context.WithCancel(context.Background())
	child1, cancelChild1 := context.WithCancel(parent1) // child1 derives from parent1.
	defer cancelChild1()                                // Always release the child.

	cancelParent1() // Cancel the parent...
	fmt.Println("cancel parent -> parent:", state(parent1), "child:", state(child1))

	// Pair 2: cancelling the child leaves the parent alive.
	parent2, cancelParent2 := context.WithCancel(context.Background())
	defer cancelParent2() // Always release the parent.
	child2, cancelChild2 := context.WithCancel(parent2)

	cancelChild2() // Cancel only the child...
	fmt.Println("cancel child  -> parent:", state(parent2), "child:", state(child2))
}
