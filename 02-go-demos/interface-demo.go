package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	City string
}

func (e *Employee) Display() {
	fmt.Printf("Employee = %v\n", e)
}

/* func printData(e Employee){
	e.Display()
} */

type Product struct {
	Id   int
	Name string
	Cost float64
}

type Displayable interface {
	Display()
}

func (p *Product) Display() {
	fmt.Printf("Product = %v\n", p)
}

func printData(obj Displayable) {
	obj.Display()
}

func main() {
	e := &Employee{100, "Magesh", "Bengaluru"}
	printData(e)

	p := &Product{500, "Pen", 10}
	printData(p)

}
