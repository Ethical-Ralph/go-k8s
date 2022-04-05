package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world\n"))
}

func startServer(errChan chan<- error) {
	errChan <- http.ListenAndServe(":3000", nil)
}

func main() {

	err := make(chan error)

	http.HandleFunc("/hello", hello)

	fmt.Println("Server started at port 3000")
	go startServer(err)

	panic(<-err)
}
