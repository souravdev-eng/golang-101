package main

import "fmt"

func main() {
	for number := 1; number <= 6; number++ {
		if number == 2 {
			continue // Skip the rest of this turn, then run the loop update.
		}
		if number == 5 {
			break // Leave this loop immediately.
		}
		fmt.Println("number:", number)
	}
}
