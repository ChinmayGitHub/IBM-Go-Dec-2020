package main

import (
	"fmt"
)

func sum(list []int, c chan int) {
	sum := 0
	for _, v := range list {
		sum += v
	}
	c <- sum
}

func main() {
	nos := []int{3, 6, 1, 4, 2, 5}
	/*
		firstChan := make(chan int)
		secondChan := make(chan int)
		first := nos[:len(nos)/2]
		second := nos[len(nos)/2:]
		go sum(first, firstChan)
		go sum(second, secondChan)
		firstTotal := <-firstChan
		secondTotal := <-secondChan
		total := firstTotal + secondTotal
		fmt.Println(total)
	*/

	c := make(chan int)
	go sum(nos[:len(nos)/2], c)
	go sum(nos[len(nos)/2:], c)
	total := <-c
	fmt.Println(total)
}
