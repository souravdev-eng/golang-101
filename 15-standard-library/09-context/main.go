package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("before cancel:", ctx.Err() == nil)

	cancel()     // Signal that work using this context should stop.
	<-ctx.Done() // The channel is closed when cancellation takes effect.
	fmt.Println("after cancel:", ctx.Err())
}
