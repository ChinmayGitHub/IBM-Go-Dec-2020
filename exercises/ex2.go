package main

import "fmt"

func main() {
	var i int = 10
	increment(&i)
	fmt.Println(i) //=> 11
}

func increment(p *int) {
	*p++
	return
}
