package main

import (
	"fmt"
	"strings"
)

type MyStr string

func (str MyStr) Uppercase() string {
	return strings.ToUpper(string(str))
}

func main() {
	fmt.Println(MyStr("This is a sample string").Uppercase())

}
