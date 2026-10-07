package main

import (
	"context"
	"fmt"
)

func main() {
	// context.Background() is the empty root context. Every context tree starts
	// from it (or from context.TODO()). It is never cancelled and carries no
	// values or deadline.
	ctx := context.Background()

	// ctx.Done() returns a channel that closes when the context is cancelled.
	// For Background that never happens, so the channel is nil. We do NOT try to
	// receive from it here: receiving from a nil channel (<-ch) blocks forever.
	fmt.Println("Done channel is nil:", ctx.Done() == nil)

	// ctx.Err() reports why the context finished. Background never finishes, so
	// its Err() is always nil.
	fmt.Println("Err is nil:", ctx.Err() == nil)
}
