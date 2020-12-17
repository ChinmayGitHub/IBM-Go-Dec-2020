package main

import "fmt"

var names = []string{"Magesh", "Mayank", "Gayathri", "Poorna", "John", "jacon", "Joseph", "Ibrahim", "Guru", "Vasu", "Bharath", "Hariprasad"}

func main() {

	//using slice
	var maxLen int
	for _, name := range names {
		if l := len(name); l > maxLen {
			maxLen = l
		}
	}
	result := make([][]string, maxLen)
	for _, name := range names {
		result[len(name)-1] = append(result[len(name)-1], name)
	}
	fmt.Println(result)

	//using map
	groupNames := make(map[int][]string)

	for _, value := range names {
		groupNames[len(value)] = append(groupNames[len(value)], value)
	}
	fmt.Println(groupNames)
}

//[["Magesh", "Mayank", "Poorna"], ["John", "Vasu", "Guru"], []]
