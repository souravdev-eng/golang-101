package main

import "fmt"

func main() {
	scores := make([]int, 0, 3) // Start empty with room for three elements.
	fmt.Println("start:", len(scores), cap(scores))
	scores = append(scores, 10, 20) // Keep the returned slice after appending.
	fmt.Println("scores:", scores)
	fmt.Println("length/capacity:", len(scores), cap(scores))
	scores = append(scores, 30)
	fmt.Println("full:", scores)
	// Appending beyond capacity allocates a larger array.
	// Go chooses the new capacity; do not depend on a particular growth size.
	scores = append(scores, 40)
	fmt.Println("grown:", scores, "length:", len(scores))
}
