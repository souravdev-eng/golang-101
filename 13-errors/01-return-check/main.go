package main

import (
	"errors" // errors.New builds an error with a message.
	"fmt"
)

func share(apples, people int) (int, error) {
	if people <= 0 {
		return 0, errors.New("people must be positive")
	}
	return apples / people, nil // nil means no error.
}

func main() {
	each, err := share(8, 2)
	if err != nil {
		fmt.Println("cannot share:", err)
		return // Stop before using an invalid result.
	}
	fmt.Println("each:", each)
	_, err = share(8, 0)
	if err != nil {
		fmt.Println("cannot share:", err)
	}
}
