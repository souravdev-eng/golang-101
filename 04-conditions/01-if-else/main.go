package main

import "fmt"

func main() {
	temperature := 24
	// Conditions do not need parentheses. Braces mark each block.
	if temperature < 15 {
		fmt.Println("wear a coat")
	} else if temperature < 25 {
		fmt.Println("wear a light jacket")
	} else {
		fmt.Println("wear a T-shirt")
	}
}
