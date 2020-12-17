package main

import (
	"fmt"
	"time"
)

func say(s string, c chan int, id int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
	}
	c <- id
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go say("Hello", c1, 1)
	go say("World", c2, 2)
	fmt.Println(<-c1)
	fmt.Println(<-c2)
	fmt.Println("End of application")
}
