package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputFile := "data.dat"
	outputFile := "new.dat"
	fileContents, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(outputFile, fileContents, 0x777)
	if err != nil {
		fmt.Println(err)
		return
	}
}
