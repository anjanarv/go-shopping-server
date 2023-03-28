package main

import (
	"net/http"
)

func main() {
	srv := NewServer()
	http.ListenAndServe(":8080", srv)
}
