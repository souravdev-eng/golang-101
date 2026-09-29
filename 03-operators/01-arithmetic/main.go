package main

import "fmt"

func main() {
	apples, baskets := 7, 2
	fmt.Println("add:", apples+baskets)
	fmt.Println("subtract:", apples-baskets)
	fmt.Println("multiply:", apples*baskets)
	fmt.Println("divide:", apples/baskets)
	fmt.Println("remainder:", apples%baskets) // % is the remainder after division.
	apples += 3                               // Shorthand for apples = apples + 3.
	apples--                                  // Subtract one.
	fmt.Println("updated:", apples)
}
