package main

import (
	"fmt"
	"slices"
)

func main() {
	prices := []int{12, 5, 9}
	sorted := slices.Clone(prices) // Make a copy before sorting.
	slices.Sort(sorted)            // Sort changes the slice passed to it.

	fmt.Println("original:", prices)
	fmt.Println("sorted:", sorted)
	fmt.Println("has 9:", slices.Contains(sorted, 9))
}
