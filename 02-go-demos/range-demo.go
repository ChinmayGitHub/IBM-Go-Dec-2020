package main

import "fmt"

func main() {
	var nos = []int{3, 1, 5, 2, 4}

	for idx, value := range nos {
		fmt.Println(idx, value)
	}

	for _, value := range nos {
		fmt.Println(value)
	}

	colors := map[string]int{"red": 1, "green": 2, "blue": 3}

	for key, value := range colors {
		fmt.Println(key, value)
	}
}
