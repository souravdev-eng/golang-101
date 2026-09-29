package main

import "fmt"

func main() {
	age := 18
	fmt.Println("equal:", age == 18) // == compares; = assigns.
	fmt.Println("different:", age != 18)
	fmt.Println("less:", age < 18, "at most:", age <= 18)
	fmt.Println("greater:", age > 18, "at least:", age >= 18)
	hasTicket, hasPass := true, false
	fmt.Println("can enter:", age >= 18 && hasTicket) // && means both are true.
	fmt.Println("has access:", hasTicket || hasPass)  // || means at least one is true.
	fmt.Println("needs ticket:", !hasTicket)          // ! flips a boolean.
}
