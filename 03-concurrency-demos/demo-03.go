package main

import (
	"fmt"
)

func sum(list []int, goId int, c chan int) {
	sum := 0
	for _, v := range list {
		sum += v
	}
	fmt.Printf("before - goId %d returning result %d \n", goId, sum)
	c <- sum
	fmt.Printf("after - goId %d returning result \n", goId)
}

func main() {
	nos := []int{3, 6, 1, 4, 2, 5}

	channel := make(chan int)
	//secondChan := make(chan int)
	first := nos[:len(nos)/2]
	second := nos[len(nos)/2:]
	go sum(first, 1, channel)
	go sum(second, 2, channel)
	fmt.Println("receiving firstTotal")
	firstTotal := <-channel
	fmt.Printf("received firstTotal - %d\n", firstTotal)
	fmt.Println("receiving secondTotal")
	secondTotal := <-channel
	fmt.Printf("received secondTotal - %d\n", secondTotal)
	total := firstTotal + secondTotal
	fmt.Println(total)

	/*
		c := make(chan int)
		go sum(nos[:len(nos)/2], c)
		go sum(nos[len(nos)/2:], c)
		total := <-c + <-c
		fmt.Println(total)
	*/
}
