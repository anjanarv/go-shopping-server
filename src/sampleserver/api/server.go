package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type Item struct {
	ID   uuid.UUID `json:"id""`
	Name string    `json:"name""`
}

type Server struct {
	*mux.Router

	shoppingItems []Item
}

func NewServer() *Server {
	s := &Server{
		Router:        mux.NewRouter(),
		shoppingItems: []Item{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/items", s.listItems()).Methods("GET")
	s.HandleFunc("/items", s.createShoppingItem()).Methods("POST")
}

func (s *Server) createShoppingItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i Item
		i.ID = uuid.New()
		s.shoppingItems = append(s.shoppingItems, i)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&i); err != nil {
			return
		}

	}
}

func (s *Server) listItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.shoppingItems); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
