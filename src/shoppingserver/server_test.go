package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShoppingServer(t *testing.T) {
	srv := NewShoppingServer(NewInMemoryShoppingStore())

	t.Run("AddShoppingItems", func(t *testing.T) {

		items := []Item{{ID: uuid.New(), Name: "test"}}

		b, err := json.Marshal(items)
		assert.NoError(t, err)

		request, _ := http.NewRequest(http.MethodPost, "/items", bytes.NewReader(b))
		recorder := httptest.NewRecorder()

		srv.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusCreated, recorder.Code)

	})

	t.Run("GetShoppingItems", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/items", nil)
		recorder := httptest.NewRecorder()

		srv.ServeHTTP(recorder, request)

		body, err := io.ReadAll(recorder.Body)
		assert.NoError(t, err)

		var items []Item
		err = json.Unmarshal(body, &items)
		assert.NoError(t, err)

		fmt.Println(items)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, 3, len(items))

	})

}
