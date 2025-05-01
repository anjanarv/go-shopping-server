# Web Server using Gorilla/mux

Build a web server using [Gorilla/mux](https://pkg.go.dev/github.com/gorilla/mux#section-readme)

#### GET `/items/` 
  should return all the items in the store

#### POST `/item` 
  should record an item for every subsequent POST

The name `mux` stands for HTTP request multiplexer. It matches the incoming requests/routes against the registered routes and calls the
corresponding handler for the route.

The implementation uses an in memory store of items.

# Testing

`httptest` package provides utilities for testing. This package is a powerful tool in the Go standard library designed specifically for testing HTTP servers and clients.
We can create mock requests and responses easily and verify if the corresponding handlers are getting invoked.

# Versioning

The entire package is versioned and release created using `gh` api on every branch merge to main.
