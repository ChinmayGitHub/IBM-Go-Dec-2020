package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4000")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		if strings.Trim(input, "\r\n") == "Q" {
			return
		}
		conn.Write([]byte(input))
	}
}
