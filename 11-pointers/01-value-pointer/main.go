package main

import "fmt"

func addToCopy(count int) {
	count++ // Only this function's local copy changes.
}

func addThroughPointer(count *int) { // *int is a pointer to an integer.
	*count = *count + 1 // * reads or writes the value at that address.
}

func main() {
	count := 3
	addToCopy(count)
	fmt.Println("after value:", count)
	addThroughPointer(&count) // & takes the address of count.
	fmt.Println("after pointer:", count)
	var missing *int // A pointer's zero value is nil: no value to point to.
	fmt.Println("missing:", missing == nil)
	// Do not dereference a nil pointer; it causes a panic.
}
