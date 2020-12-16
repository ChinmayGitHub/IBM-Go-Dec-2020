package main

import "fmt" //importing dependant packages

//package level variables

//declaring 1 variable
var id int

//declaring multiple variables
/*
var (
	name string
	age int
	location string
)
*/

/*
var (
	name, location string
	age int
)
*/

//variables with initilizers
/*
var (
	name string = "Magesh Kuppan"
	age int = 35
	location string = "Bengaluru"
)
*/

//data types can be omitted when initialized
/*
var (
	name = "Magesh Kuppan"
	age = 35
	location = "Bengaluru"
)
*/

//
/*
var (
	name, age, location = "Magesh Kuppan", 35, "Bengaluru"
)
*/

func init() { // function executed when the package is initialized

}

func main() {

	const pi = 3.14

	const (
		red   = 1
		green = 2
		blue  = 3
	)

	name := "Magesh Kuppan"
	age := 45
	location := "Bengaluru"

	fmt.Println(pi)
	s := fmt.Sprintf("%s (%d) %s\n", name, age, location)
	fmt.Printf("%T (%T) %T\n", name, age, location)
	fmt.Println(s)

	//type conversion
	/*
		var v1 int32 = 100
		var v2 int64
		var f float64
		v2 = int64(v1)
		f = float64(v1)
		fmt.Printf("%v, %v\n", v2, f)
	*/

	var v1 float64 = 20000.9999
	var v2 int8
	v2 = int8(v1)
	fmt.Printf("%v\n", v2)
}

//interfaces

//struct

//functions
