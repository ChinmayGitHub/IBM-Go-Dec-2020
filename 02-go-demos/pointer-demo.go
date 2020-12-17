package main

import "fmt"

func main() {
	var p *int
	fmt.Println(p)

	var i int
	p = &i

	fmt.Println(p)
	fmt.Println(*p) //dereferencing the value using the address

	*p = 10
	fmt.Println("i = ", i)
}
