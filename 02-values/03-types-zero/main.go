package main

import "fmt"

func main() {
	var count int     // Whole numbers start at 0.
	var price float64 // Decimal numbers start at 0.
	var ready bool    // A boolean is true or false; its zero value is false.
	var name string   // Strings start empty.
	// %q puts quotes around text so an empty string is visible.
	fmt.Printf("zero: %d %.1f %t %q\n", count, price, ready, name)
	count, price, ready, name = 2, 3.5, true, "tea"
	// %T displays the type of a value.
	fmt.Printf("types: %T %T %T %T\n", count, price, ready, name)
	fmt.Println(count, price, ready, name)
}
