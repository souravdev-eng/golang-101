package main

import "fmt"

func main() {
	message := "outside"
	fmt.Println(message)
	if true {
		message := "inside" // A new variable shadows the outer one in this block.
		fmt.Println(message)
	}
	fmt.Println(message) // The outer variable has not changed.
	if true {
		message = "updated" // Assignment uses the existing outer variable.
	}
	fmt.Println(message)
	// Prefer different names when shadowing would make the code confusing.
}
