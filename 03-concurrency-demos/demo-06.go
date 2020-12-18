package main

import "fmt"

func main() {
	c := make(chan int, 3)
	go func() {
		c <- 10
		c <- 20
		c <- 30
		close(c)
	}()

	/*
		fmt.Println(<-c)
		fmt.Println(<-c)
		fmt.Println(<-c)
	*/

	for i := range c {
		fmt.Println(i)
	}
}
