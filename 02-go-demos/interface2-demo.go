package main

import "fmt"

var x = 10
var str = "MyString"

type Product struct {
	name string
	cost float64
}

type Any interface{}

func main() {
	var myVar Any
	myVar = 10
	fmt.Printf("myVar has the value: %v\n", myVar)
	myVar = str
	fmt.Printf("myVar has the value: %v\n", myVar)
	prod := &Product{name: "Pen", cost: 50}
	myVar = prod
	fmt.Printf("myVar has the value: %v\n", myVar)
}
