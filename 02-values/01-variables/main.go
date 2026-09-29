package main

import "fmt"

func main() {
	var apples int = 3       // var can name the type explicitly.
	var shop = "Corner shop" // Go infers string from this value.
	baskets := 2             // := declares a new local variable with an inferred type.
	apples = 5               // = changes an existing variable.
	baskets++                // Add one to the existing value.
	fmt.Println(shop)
	fmt.Println("apples:", apples, "baskets:", baskets)
}
