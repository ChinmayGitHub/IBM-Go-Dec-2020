package main

import (
	"bufio"
	"os"
)

func main() {
	outputFile, _ := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0644)
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	for i := 0; i < 100; i++ {
		outputWriter.WriteString("This is a line \n")
	}
	outputWriter.Flush()

}
