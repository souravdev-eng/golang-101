package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "pear", "plum"}
	fmt.Println("first:", fruits[0], "length:", len(fruits))
	middle := fruits[1:3] // Include index 1, stop before index 3.
	fmt.Println("middle:", middle)
	middle[0] = "orange" // Both slices refer to the same underlying array.
	fmt.Println("fruits:", fruits)
	fmt.Println("first two:", fruits[:2]) // An omitted start means zero.
	fmt.Println("last two:", fruits[2:])  // An omitted end means len(fruits).
	// Valid element indexes are 0 through len(fruits)-1.
}
