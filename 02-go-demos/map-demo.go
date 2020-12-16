package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["foo"] = 10
	m["bar"] = 20

	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)

	colors := map[string]int{"red": 1, "green": 2, "blue": 3}
	fmt.Println(colors["green"])
	fmt.Println(len(colors))

	val, found := colors["purple"]
	fmt.Println(val, found)

	if _, exists := colors["purple"]; exists {
		fmt.Println("purple exists")
	} else {
		fmt.Println("purple doesnot exist")
	}

}
