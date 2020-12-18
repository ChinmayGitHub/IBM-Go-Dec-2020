package main

import (
	"fmt"
	"net/http"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello %s ", req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", helloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
