package main

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

// ShoppingStore interface for Get/Add Items
type ShoppingStore interface {
	GetItems() ([]byte, error)
	AddItems(byte []byte) error
}

type Item struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type ShoppingServer struct {
	*mux.Router

	store ShoppingStore
}

// NewShoppingServer creates a new server with store and router
func NewShoppingServer(store ShoppingStore) *ShoppingServer {
	s := &ShoppingServer{
		Router: mux.NewRouter(),
		store:  store,
	}
	s.routes()
	return s
}

// routes is for defining the routes for the web server
func (s *ShoppingServer) routes() {
	s.HandleFunc("/items", s.AddItems()).Methods("POST")
	s.HandleFunc("/items", s.GetItems()).Methods("GET")
}

func (s *ShoppingServer) AddItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("Adding Items")
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		if r.Body == nil {
			log.Println("Request body is nil")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// read request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Failed to read request body")
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		defer r.Body.Close() // nolint

		// add items to store
		err = s.store.AddItems(body)
		if err != nil {
			log.Println("Failed to get shopping items")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		log.Println("Successfully added Items")

		w.WriteHeader(http.StatusCreated)
	}
}

func (s *ShoppingServer) GetItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println("Getting Items")
		if r.Method != http.MethodGet {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		// get items from store
		items, err := s.store.GetItems()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("Successfully received Items:", string(items))
		// write into response stream
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(items)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
