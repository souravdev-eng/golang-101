package main

import "fmt"

type Book struct { // type introduces a named type.
	Title string
	Pages int
}

func main() {
	book := Book{Title: "Go Basics", Pages: 80}
	fmt.Println(book.Title, book.Pages) // A dot selects a field.
	book.Pages = 96
	fmt.Println("updated pages:", book.Pages)
	var empty Book // Each field starts with its zero value.
	fmt.Printf("zero book: %q %d\n", empty.Title, empty.Pages)
}
