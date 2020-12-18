package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	inputFile, inputErr := os.Open("employee.dat")
	if inputErr != nil {
		fmt.Println(inputErr)
		return
	}
	defer inputFile.Close()
	inputReader := csv.NewReader(inputFile)
	sum := float64(0)
	count := 0
	for {
		data, err := inputReader.Read()
		if err == io.EOF {
			break
		}
		sal, _ := strconv.ParseFloat(data[2], 64)
		sum += sal
		count++
	}
	fmt.Println(sum / float64(count))
}
