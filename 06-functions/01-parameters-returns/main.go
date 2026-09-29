package main

import "fmt"

func greet(name string) { // A parameter names the input and its type.
	fmt.Println("Hello,", name)
}

func total(price int, quantity int) int { // The final int is the return type.
	return price * quantity
}

func main() {
	greet("Mira")
	cost := total(3, 4)
	fmt.Println("cost:", cost)
}
