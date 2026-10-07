package main

import (
	"context"
	"fmt"
	"time"
)

// describe waits for a context to finish, then reports why it finished.
func describe(ctx context.Context, label string) {
	<-ctx.Done() // Block until the context is done. Done() is a channel.

	// After Done() fires, Err() says WHY. The two built-in reasons are
	// context.Canceled and context.DeadlineExceeded. switch branches on which
	// one we got.
	switch ctx.Err() {
	case context.Canceled:
		fmt.Println(label, "was cancelled by a caller")
	case context.DeadlineExceeded:
		fmt.Println(label, "ran out of time")
	}
}

func main() {
	// Case 1: cancelled explicitly -> ctx.Err() is context.Canceled.
	ctx1, cancel1 := context.WithCancel(context.Background())
	cancel1() // Cancel right away so Done() is already closed.
	describe(ctx1, "manual:")

	// Case 2: timed out -> ctx.Err() is context.DeadlineExceeded.
	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel2() // Still release resources even though it times out.
	describe(ctx2, "timeout:")
}
