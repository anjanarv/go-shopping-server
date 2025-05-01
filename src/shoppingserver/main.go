package main

import (
	"net/http"
)

func main() {
	srv := NewShoppingServer(NewInMemoryShoppingStore())

	// create a webserver in Go
	// the below call creates a goroutine for every request and runs it against the handler(2nd param)
	err := http.ListenAndServe(":8080", srv)
	if err != nil {
		return
	}
}
