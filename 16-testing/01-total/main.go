package main

import "fmt"

// Total calculates the cost of a whole-number price and quantity.
func Total(price, quantity int) int {
	return price * quantity
}

func main() {
	fmt.Println("cost:", Total(3, 4))
}
