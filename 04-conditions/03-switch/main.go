package main

import "fmt"

func main() {
	day := "Saturday"
	switch day {
	case "Saturday", "Sunday": // Several values can share one branch.
		fmt.Println("weekend")
	case "Monday":
		fmt.Println("start of work week")
	default:
		fmt.Println("weekday")
	}
	// Go finishes a matched case automatically; no break is needed.
	score := 82
	switch { // With no expression, the first true condition wins.
	case score >= 90:
		fmt.Println("excellent")
	case score >= 70:
		fmt.Println("good progress")
	default:
		fmt.Println("keep practicing")
	}
}
