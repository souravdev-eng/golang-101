package main

import (
	"fmt"
	"strconv" // strconv converts between text and numeric values.
)

func printQuantity(text string) {
	quantity, err := strconv.Atoi(text) // Atoi parses an integer from a string.
	if err != nil {
		// Keep the lesson's output simple; err also carries diagnostic details.
		fmt.Printf("invalid quantity: %q\n", text)
		return
	}
	fmt.Println("quantity:", quantity)
}

func main() {
	printQuantity("12")
	printQuantity("many")
}
