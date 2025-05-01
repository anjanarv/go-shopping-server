package main

import (
	"encoding/json"
	"github.com/google/uuid"
)

func NewInMemoryShoppingStore() *InMemoryShoppingStore {
	return &InMemoryShoppingStore{
		[]Item{
			{uuid.New(), "test1"},
			{uuid.New(), "test2"},
		},
	}
}

type InMemoryShoppingStore struct {
	items []Item
}

func (i *InMemoryShoppingStore) AddItems(b []byte) error {
	var items []Item

	err := json.Unmarshal(b, &items)
	if err != nil {
		return err
	}

	i.items = append(i.items, items...)

	return nil
}

func (i *InMemoryShoppingStore) GetItems() ([]byte, error) {
	b, err := json.Marshal(i.items)
	if err != nil {
		return nil, err
	}

	return b, nil
}
