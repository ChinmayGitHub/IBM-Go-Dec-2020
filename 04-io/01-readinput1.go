package main

import "fmt"

var (
	firstname string
	lastname  string
	s         string
	i         int
	f         float64
)

func main() {
	fmt.Println("Enter your first and last name")
	fmt.Scanln(&firstname, &lastname)
	fmt.Printf("Hi %s %s!\n", firstname, lastname)
	input := "100.5 / 2000 / Magesh"
	fmt.Sscanf(input, "%f / %d / %s", &f, &i, &s)
	fmt.Println(f, i, s)
}
