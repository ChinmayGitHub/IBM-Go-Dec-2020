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

var (
	name, age, location = "Magesh Kuppan", 35, "Bengaluru"
)

func init() { // function executed when the package is initialized

}

func main() {
	fmt.Println("Welcome to Golang!")
}

//interfaces

//struct

//functions
