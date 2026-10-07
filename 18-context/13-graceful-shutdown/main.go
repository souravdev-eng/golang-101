package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// A long-running program should stop on its own terms when the operating system
// asks it to quit (Ctrl-C sends SIGINT). signal.NotifyContext turns that signal
// into a cancelled context: instead of dying mid-request, the program notices
// ctx.Done(), finishes the work already in flight, and exits cleanly.
//
// signal.NotifyContext would normally wait for a real Ctrl-C. To keep the
// example deterministic and hands-free, a goroutine sends the process SIGINT
// itself after a short delay, standing in for the human pressing Ctrl-C.

func main() {
	// NotifyContext returns a context that is cancelled when SIGINT arrives, and
	// a stop function that releases the signal handler.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Stand in for the human: raise SIGINT at this process after a brief pause,
	// once main is already waiting below.
	go func() {
		time.Sleep(10 * time.Millisecond)
		_ = syscall.Kill(os.Getpid(), syscall.SIGINT)
	}()

	fmt.Println("server: running, waiting for shutdown signal")
	<-ctx.Done() // Blocks until the signal cancels the context.

	// We were asked to stop. Drain the in-flight work rather than dropping it.
	fmt.Println("server: shutdown signal received, draining in-flight work")
	for job := 1; job <= 2; job++ {
		time.Sleep(5 * time.Millisecond) // Pretend each job takes a moment.
		fmt.Printf("server: finished job %d\n", job)
	}
	fmt.Println("server: exited cleanly")
}
