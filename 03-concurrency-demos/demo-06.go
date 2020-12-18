package main

import "fmt"

func main() {
	c := make(chan int, 2)
	go func() {

		c <- 10
		//fmt.Println("Writing 10")

		c <- 20
		//fmt.Println("Writing 20")

		c <- 30
		//fmt.Println("Writing 30")

		c <- 40
		fmt.Println("Writing 40")

		c <- 50
		fmt.Println("Writing 50")
		//close(c)
	}()

	/*
		fmt.Println(<-c)
		fmt.Println(<-c)
		fmt.Println(<-c)
	*/

	/* for i := range c {
		fmt.Println(i)
	} */

	/* for i := 0; i < 3; i++ {
		fmt.Println(<-c)
	} */

	<-c
	//<-c
	//<-c
	fmt.Println("Done")

}
