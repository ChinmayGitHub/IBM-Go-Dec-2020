package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1000 * time.Millisecond)
	end := time.After(5000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Print("Tick")
		case <-end:
			fmt.Println("Done")
			return
		default:
			fmt.Print(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
