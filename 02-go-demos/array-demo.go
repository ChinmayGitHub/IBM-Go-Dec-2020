package main

import "fmt"

func main() {
	var nos [5]int
	fmt.Println(nos)

	nos[3] = 100

	newNos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(newNos)
	fmt.Println(newNos[:2])
}
