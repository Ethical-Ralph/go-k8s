package main

import (
	"fmt"
	"net/http"
)

func dummy(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dummy service\n"))
}

func startServer(errChan chan<- error) {
	errChan <- http.ListenAndServe(":3000", nil)
}

func main() {

	err := make(chan error)

	http.HandleFunc("/dummy", dummy)

	fmt.Println("Server started at port 3000")
	go startServer(err)

	panic(<-err)
}
