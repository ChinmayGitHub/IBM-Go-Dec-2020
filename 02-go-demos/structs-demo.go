package main

import "fmt"

type Person struct {
	Id         int
	Name       string
	IsEmployed bool
	City       string
}

type Employee struct {
	Person
	Org string
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

//as methods
func (p Person) display(format string) string {
	//fmt.Printf(format, p.Id, p.Name, p.IsEmployed)
	return fmt.Sprintf(format, p.Id, p.Name, p.IsEmployed)
}

func (p *Person) toggleEmploymentStatus() {
	p.IsEmployed = !p.IsEmployed
}

func main() {
	//Composite Literal syntax
	/* var p = Person{Id: 100, Name: "Magesh", IsEmployed: true, City: "Bangalore"} */
	//p := &Person{100, "Magesh", true, "Bangalore"}

	//using the 'new' keyword
	p := new(Person)
	p.Id = 100
	p.Name = "Magesh"
	p.IsEmployed = true
	p.City = "Bangalore"
	/*
		display(p)
		toggleEmploymentStatus(&p)
		display(p)
	*/
	fmt.Println(p.display("Id=%d, Name=%s, IsEmployed=%v"))
	p.toggleEmploymentStatus()
	fmt.Println(p.display("Id=%d, Name=%s, IsEmployed=%v"))

	//using composite Literal syntax
	//e := Employee{Person{100, "Magesh", true, "Bangalore"}, "IBM"}

	//using 'new' keyword
	e := new(Employee) // equivalent to  e := &Employee{Person{100, "Magesh", true, "Bangalore"}, "IBM"}
	e.Id = 100
	e.Name = "Magesh"
	e.IsEmployed = true
	e.City = "Bangalore"
	e.Org = "IBM"
	fmt.Println(e)
	fmt.Println(e.Id, e.Person.Id)

	fmt.Println(e.display("Id=%d, Name=%s, IsEmployed=%v"))
	e.toggleEmploymentStatus()
	fmt.Println(e.display("Id=%d, Name=%s, IsEmployed=%v"))
}
