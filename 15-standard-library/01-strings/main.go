package main

import (
	"fmt"
	"strings" // Standard-library helpers for text.
)

func main() {
	input := "  Tea,Milk,Bread  "
	clean := strings.TrimSpace(input) // Remove surrounding whitespace.
	clean = strings.ToLower(clean)
	items := strings.Split(clean, ",") // Split returns a slice of strings.
	fmt.Println("items:", items)
	fmt.Println("has milk:", strings.Contains(clean, "milk"))
	fmt.Println("display:", strings.Join(items, " | "))
}
