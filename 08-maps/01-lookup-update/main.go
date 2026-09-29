package main

import "fmt"

func main() {
	stock := map[string]int{"apples": 3, "pears": 2}
	fmt.Println("apples:", stock["apples"])
	stock["apples"] = 5  // Update an existing key.
	stock["oranges"] = 4 // Insert a new key.
	delete(stock, "pears")
	count, exists := stock["pears"] // The second result says whether the key exists.
	fmt.Println("pears:", count, "exists:", exists)
	stock["bananas"] = 0
	count, exists = stock["bananas"]
	fmt.Println("bananas:", count, "exists:", exists)
	fmt.Println("keys:", len(stock))
	// Missing integer values return 0, so use exists to tell missing from zero.
}
