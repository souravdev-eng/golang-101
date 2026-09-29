package main

import "fmt"

func splitApples(apples, people int) (int, int) {
	// This lesson assumes people is positive; errors are covered in topic 13.
	return apples / people, apples % people
}

func main() {
	each, left := splitApples(7, 3) // Receive the results in their declared order.
	fmt.Println("each:", each, "left:", left)
}
