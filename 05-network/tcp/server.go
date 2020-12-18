package main

import (
	"fmt"
	"net"
)

func handleRequest(conn net.Conn) {
	for {
		buff := make([]byte, 512)
		_, err := conn.Read(buff)
		if err != nil {
			fmt.Println("error reading ", err.Error())
			break
		}
		fmt.Printf("Received data : %v", string(buff))
	}
}

func main() {
	fmt.Println("Starting the server...")
	listener, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting connection", err.Error())
			return
		}
		go handleRequest(conn)
	}

}
