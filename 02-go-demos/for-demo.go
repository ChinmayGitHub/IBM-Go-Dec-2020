package main

import "fmt"

func main() {

	i := 1

	//using 'for' as a 'while' construct
	for i <= 3 {
		i += 1
	}

	//traditional for loop
	for j := 10; j <= 20; j++ {
		fmt.Println(j)
	}

	//using 'while(true)' loop
	for {
		fmt.Println("Inside an possibly indefinite loop")
		break // continue
	}

}
