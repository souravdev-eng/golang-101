package main

import "fmt"

func main() {
	fmt.Print("Shopping: ")  // Print leaves the cursor on the same line.
	fmt.Println("apples", 3) // Println separates values with spaces.
	// %s is a string, %d is an integer, and %.2f shows two decimal places.
	// Printf needs an explicit \n escape for a newline.
	fmt.Printf("%s: %d items, %.2f each\n", "apples", 3, 1.5)
}
