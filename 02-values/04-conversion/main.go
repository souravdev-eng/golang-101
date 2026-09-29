package main

import "fmt"

func main() {
	total := 7
	people := 2
	// Integer division drops the fractional part.
	fmt.Println("whole share:", total/people)
	// Convert both operands before dividing to keep the fraction.
	share := float64(total) / float64(people)
	fmt.Printf("decimal share: %.1f\n", share)
	fmt.Println("back to int:", int(share)) // Conversion drops the fraction.
}
