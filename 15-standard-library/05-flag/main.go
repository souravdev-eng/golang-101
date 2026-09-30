package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "Gopher", "name to greet") // Returns a *string.
	times := flag.Int("times", 2, "number of greetings")
	flag.Parse() // Read flags passed after the go run command.

	for i := 0; i < *times; i++ {
		fmt.Println("Hello,", *name)
	}
}
