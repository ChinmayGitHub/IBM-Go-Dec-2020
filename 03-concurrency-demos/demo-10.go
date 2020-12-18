package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://info.cern.ch/hypertext/WWW/TheProject.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
