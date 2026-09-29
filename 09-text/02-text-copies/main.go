package main

import "fmt"

func main() {
	greeting := "hello"
	letters := []byte(greeting) // byte is an alias for uint8, a number from 0 to 255.
	letters[0] = 'H'            // Single quotes describe one rune; H also fits in one byte.
	fmt.Println("original:", greeting)
	fmt.Println("edited:", string(letters))
	word := []rune("café") // Use runes to edit whole Unicode code points.
	word[3] = 's'
	fmt.Println("rune edit:", string(word))
	fmt.Println("joined:", string(letters)+" Go")
}
