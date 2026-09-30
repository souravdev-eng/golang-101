package main

import (
	"fmt"
	"time"
)

func main() {
	// Go uses this reference date to describe the input format.
	date, err := time.Parse("2006-01-02", "2024-01-01")
	if err != nil {
		fmt.Println("invalid date:", err)
		return
	}

	fmt.Println("date:", date.Format("Mon, 02 Jan 2006"))
	fmt.Println("one week later:", date.AddDate(0, 0, 7).Format("2006-01-02"))
}
