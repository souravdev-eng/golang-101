package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "pear"} // []string makes a slice of strings.
	for index, fruit := range fruits {            // range gives an index and a value.
		fmt.Println(index, fruit) // Indexes start at zero.
	}
	for _, fruit := range fruits { // _ discards the index when it is not needed.
		fmt.Println("packed:", fruit)
	}
}
