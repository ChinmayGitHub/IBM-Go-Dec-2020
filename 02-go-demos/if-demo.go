package main

import "fmt"

func main() {

	if n := 7; n%2 == 0 {
		fmt.Printf("%d is even\n", n)
	} else {
		fmt.Printf("%d is odd\n", n)
	}
}
