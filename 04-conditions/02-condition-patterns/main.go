package main

import "fmt"

func main() {
	age, hasTicket := 20, true
	if age >= 18 && hasTicket {
		fmt.Println("adult entry allowed")
	}
	if hasTicket {
		if age < 25 { // This condition is checked only when there is a ticket.
			fmt.Println("young visitor discount")
		}
	}
	// remaining exists only in this if and its else block.
	if remaining := 10 - 3; remaining > 0 {
		fmt.Println("seats left:", remaining)
	} else {
		fmt.Println("sold out")
	}
}
