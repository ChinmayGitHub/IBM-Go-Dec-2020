package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Enter some data :")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Println("The input is ", input)
	}

}
