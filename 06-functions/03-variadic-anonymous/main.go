package main

import "fmt"

func sum(numbers ...int) int { // ... lets callers supply zero or more integers.
	total := 0
	for _, number := range numbers { // Inside sum, numbers is a slice.
		total += number
	}
	return total
}

func main() {
	fmt.Println("sum:", sum(2, 3, 4))
	fmt.Println("empty sum:", sum())
	double := func(number int) int { // This function has no declared name.
		return number * 2
	}
	fmt.Println("double:", double(5))
}
