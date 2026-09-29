package main

import "fmt"

func sendLunches(lunches chan string) { // chan string carries string values.
	lunches <- "sandwich" // Send a value; an unbuffered send waits for a receiver.
	lunches <- "fruit"
	close(lunches) // The sender closes the channel after its final send.
}

func main() {
	lunches := make(chan string) // This channel has no buffer.
	go sendLunches(lunches)
	for lunch := range lunches { // Receive until the channel is closed and drained.
		fmt.Println("received:", lunch)
	}
	fmt.Println("all lunches received")
	// Receiving without a possible sender would block forever.
	// Do not send to a closed channel; closing is not a way to cancel a sender.
}
