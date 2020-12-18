package main

import (
	"fmt"
	"time"
)

func fibonacci(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			fmt.Println(<-c)
		}
	}()
	/*
		go func() {
			var input string
			fmt.Scanln(&input)
			quit <- 0
		}()
		fibonacci(c, quit)
	*/

	go fibonacci(c, quit)
	var input string
	fmt.Scanln(&input)

}
