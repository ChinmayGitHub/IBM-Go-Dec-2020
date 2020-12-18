package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, inputError := os.Open("data.dat")
	defer inputFile.Close()
	if inputError != nil {
		fmt.Println(inputError)
		return
	}
	inputReader := bufio.NewReader(inputFile)
	for {
		line, err := inputReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}

}
