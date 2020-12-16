package main

import (
	"fmt"
	"time"
)

func main() {

	no := 2
	switch no {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its a weekend!")
	default:
		fmt.Println("Its a weekday!")
	}

	tyme := time.Now()
	switch {
	case tyme.Hour() < 12:
		fmt.Println("Good Morning")
	default:
		fmt.Println("Good Evening")
	}

}
