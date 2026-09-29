package main

import "fmt"

const daysPerWeek = 7 // A constant cannot be assigned a new value.

func main() {
	const weeks = 2 // Constants can also be declared inside a function.
	fmt.Println("days:", weeks*daysPerWeek)
	// Unlike variables, constants do not use :=.
}
