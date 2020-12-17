package main

import "fmt"

type Person struct {
	Id         int
	Name       string
	IsEmployed bool
	City       string
}

func main() {
	var p = Person{Id: 100, Name: "Magesh", IsEmployed: true, City: "Bangalore"}
	fmt.Println(p)
}
