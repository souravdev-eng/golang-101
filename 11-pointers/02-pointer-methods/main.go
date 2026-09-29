package main

import "fmt"

type Counter struct {
	Value int
}

func (counter Counter) IncrementCopy() {
	counter.Value++ // The receiver is a copy.
}

func (counter *Counter) Increment() {
	counter.Value++ // Go lets us use . on a struct pointer without writing (*counter).
}

func main() {
	counter := Counter{Value: 3}
	counter.IncrementCopy()
	fmt.Println("after value method:", counter.Value)
	counter.Increment() // Go takes the address of this variable automatically.
	fmt.Println("after pointer method:", counter.Value)
}
