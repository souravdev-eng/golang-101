package main

import "fmt"

func main() {
	scores := [3]int{10, 20, 30} // Exactly three integers.
	fmt.Println("scores:", scores)
	fmt.Println("first:", scores[0], "length:", len(scores))
	copyOfScores := scores // Assigning an array copies all its elements.
	copyOfScores[0] = 99
	fmt.Println("original:", scores)
	fmt.Println("copy:", copyOfScores)
	var empty [2]int // Each element gets its type's zero value.
	fmt.Println("zero array:", empty)
}
