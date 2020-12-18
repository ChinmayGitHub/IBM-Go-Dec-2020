package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	response := make(chan *http.Response, 1)
	//errors := make(chan *error)
	go func() {
		resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
		if err != nil {
			fmt.Println("Something went wrong")
		}
		response <- resp
	}()
	for {
		select {
		case res := <-response:
			io.Copy(os.Stdout, res.Body)
			return
		case <-time.After(100 * time.Millisecond):
			fmt.Println("timeout!")
			return
		}
	}

}
