package main

import "fmt"

func pack() {
	fmt.Println("start packing")
	defer fmt.Println("close box") // Runs last: deferred calls run in reverse order.
	defer fmt.Println("put lid on")
	label := "tea"
	defer fmt.Println("label:", label) // Arguments are evaluated at the defer line.
	label = "coffee"
	fmt.Println("current label:", label)
	fmt.Println("packing done")
	// All three deferred calls run as pack returns, before main continues.
}

func main() {
	pack()
	fmt.Println("back in main")
}
