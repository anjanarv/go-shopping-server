package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testServer *ShoppingServer

func setup() {
	if testServer == nil {
		testServer = NewShoppingServer(NewInMemoryShoppingStore())
	}
}

func TestInMemoryShoppingStore(t *testing.T) {
	setup() // create new server

	t.Run("Add items", func(t *testing.T) {
		items := []Item{{ID: uuid.New(), Name: "test"}}

		b, err := json.Marshal(items)
		assert.NoError(t, err)

		err = testServer.store.AddItems(b)
		assert.NoError(t, err)

	})

	t.Run("Get items", func(t *testing.T) {

		items, err := testServer.store.GetItems()
		assert.NoError(t, err)

		var itemsReceived []Item
		err = json.Unmarshal(items, &itemsReceived)
		assert.NoError(t, err)

		assert.NotNil(t, itemsReceived)
		assert.Equal(t, 3, len(itemsReceived))
	})
}
