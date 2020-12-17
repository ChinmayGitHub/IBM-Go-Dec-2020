package main

import "fmt"

var x = 10
var str = "MyString"

type Product struct {
	name string
	cost float64
}

type Any interface{}

func printType(myVar Any) {
	switch t := myVar.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type bool %T\n", t)
	case *Product:
		fmt.Printf("Type pointer to product %T\n", *t)
	default:
		fmt.Printf("unexpected type %T\n", t)
	}
}

func main() {
	var myVar Any
	myVar = 10
	fmt.Printf("myVar has the value: %v\n", myVar)
	printType(myVar)
	myVar = str
	fmt.Printf("myVar has the value: %v\n", myVar)
	printType(myVar)
	prod := &Product{name: "Pen", cost: 50}
	myVar = prod
	fmt.Printf("myVar has the value: %v\n", myVar)
	printType(myVar)
	printType(float64(100))
}
