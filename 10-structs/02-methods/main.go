package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func (rectangle Rectangle) Area() int { // rectangle is a value receiver: a copy.
	return rectangle.Width * rectangle.Height
}

func main() {
	garden := Rectangle{Width: 4, Height: 3}
	fmt.Println("area:", garden.Area())
	// Calling a method uses a dot, just like selecting a field.
}
