package main

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Login endpoint\n"))
}

func startServer(errChan chan<- error) {
	errChan <- http.ListenAndServe(":3000", nil)
}

func main() {

	err := make(chan error)

	http.HandleFunc("/login", login)

	fmt.Println("Server started at port 3000")
	go startServer(err)

	panic(<-err)
}
