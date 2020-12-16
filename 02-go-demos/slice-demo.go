package main

import "fmt"

func main() {
	nos := make([]int, 3)
	fmt.Println(nos)
	nos[0] = 10
	nos[1] = 20
	nos[2] = 30
	fmt.Println(nos)

	nos = append(nos, 40)
	nos = append(nos, 50, 60, 70)
	fmt.Println(nos)

	fmt.Println(nos[2:])
	fmt.Println(nos[:5])
	fmt.Println(nos[2:5])

	fmt.Println(len(nos))
}
