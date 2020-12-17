package main

import "fmt"

type Person struct {
	Id         int
	Name       string
	IsEmployed bool
	City       string
}

//as functions
/*
func  display(p Person){
	fmt.Println(p)
	return
}

func  toggleEmploymentStatus (p *Person) {
	p.IsEmployed = !p.IsEmployed
}
*/

func (p Person) display() {
	fmt.Println(p)
	return
}

func (p *Person) toggleEmploymentStatus() {
	p.IsEmployed = !p.IsEmployed
}

func main() {
	/* var p = Person{Id: 100, Name: "Magesh", IsEmployed: true, City: "Bangalore"} */
	p := &Person{100, "Magesh", true, "Bangalore"}
	/*
		display(p)
		toggleEmploymentStatus(&p)
		display(p)
	*/
	p.display()
	p.toggleEmploymentStatus()
	p.display()
}
