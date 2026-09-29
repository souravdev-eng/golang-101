package main

import (
	"fmt"
	"sort" // sort.Strings puts strings in ascending order.
)

func main() {
	stock := make(map[string]int) // Allocate an empty writable map.
	// A map declared only with var would be nil; initialize it before adding keys.
	stock["pears"] = 2
	stock["apples"] = 3
	keys := make([]string, 0, len(stock))
	for key := range stock { // Request just the keys.
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, stock[key])
	}
}
