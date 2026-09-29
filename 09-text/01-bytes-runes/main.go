package main

import "fmt"

func main() {
	word := "café"
	fmt.Println("bytes:", len(word))         // é takes two bytes in UTF-8.
	fmt.Println("first byte:", word[0])      // Indexing a string returns a byte.
	fmt.Println("runes:", len([]rune(word))) // A rune represents a Unicode code point.
	for offset, letter := range word {
		// The offset is a byte position, not a rune count. %c displays a rune.
		fmt.Printf("byte %d: %c\n", offset, letter)
	}
	// A visible character can contain several code points, such as an added accent.
}
