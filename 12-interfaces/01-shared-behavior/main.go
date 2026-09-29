package main

import "fmt"

type Greeter interface {
	Greeting() string // The interface specifies behavior, not stored fields.
}

type Person struct {
	Name string
}

func (person Person) Greeting() string {
	return "Hello, " + person.Name
}

type Robot struct{}

func (robot Robot) Greeting() string {
	return "Beep, hello"
}

func welcome(guest Greeter) {
	fmt.Println(guest.Greeting()) // The caller does not need the concrete type.
}

func main() {
	// No implements keyword: matching methods are enough.
	welcome(Person{Name: "Mira"})
	welcome(Robot{})
}
