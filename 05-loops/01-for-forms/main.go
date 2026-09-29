package main

import "fmt"

func main() {
	// Initialize once, check before each turn, update after each turn.
	for step := 1; step <= 3; step++ {
		fmt.Println("step:", step)
	}
	remaining := 2
	for remaining > 0 { // Go uses for for while-style loops too.
		fmt.Println("remaining:", remaining)
		remaining--
	}
	attempt := 0
	for { // No condition means repeat until something exits the loop.
		attempt++
		fmt.Println("attempt:", attempt)
		if attempt == 2 {
			break
		}
	}
}
