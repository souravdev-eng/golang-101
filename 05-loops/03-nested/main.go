package main

import "fmt"

func main() {
	for row := 1; row <= 2; row++ {
		for seat := 1; seat <= 3; seat++ {
			fmt.Printf("row %d seat %d\n", row, seat)
		}
		fmt.Println("---")
	}
	// A break inside the inner loop would only leave that inner loop.
}
