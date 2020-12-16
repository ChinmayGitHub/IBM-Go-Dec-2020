package main

import "fmt"

/* func add(x int, y int) int {
	return x + y
} */

func add(x, y int) int {
	return x + y
}

//variadic functions
func sum(nos ...int) int {
	fmt.Println(nos)
	result := 0
	for index := 0; index < len(nos); index++ {
		result += nos[index]
	}
	return result
}

func divide(x int, y int) (int, int) {
	result := int(x / y)
	remainder := x % y
	return result, remainder
}

func genEvenNumbers() func() int {
	n := 0
	return func() int {
		n += 2
		return n
	}
}

func main() {
	fmt.Println(add(100, 200))
	result, remainder := divide(10, 3)
	fmt.Println(result, remainder)
	fmt.Println(sum(10, 20, 30, 40))

	getEvenNo := genEvenNumbers()
	fmt.Println(getEvenNo())
	fmt.Println(getEvenNo())
	fmt.Println(getEvenNo())
	fmt.Println(getEvenNo())
}
