package store_test

import (
	"fmt"
	"testing"

	"github.com/chrisdoherty4/rememberme/internal/pkg/store"
	"github.com/chrisdoherty4/rememberme/internal/pkg/todo"
	"github.com/stretchr/testify/assert"
)

func TestMemoryStore(t *testing.T) {
	store := store.NewMemoryStore()
	assert.Zero(t, store.Size(), "List should be initialized with 0 items")

	store.Add(todo.NewItem("Walk dog"))
	assert.Equal(t, store.Size(), 1, "List should have 1 item in")
}

func TestAddMultipleItems(t *testing.T) {
	items := []struct {
		title       string
		description string
	}{
		{"Grocery shopping", "Apples, Bananas, Oranges"},
		{"Clean car", "Clean car"},
		{"Walk dog", ""},
	}

	store := store.NewMemoryStore()

	for i, data := range items {
		item := todo.NewItem(data.title)
		item.SetDescription(data.description)

		store.Add(item)
		assert.Equal(t, store.Size(), i+1, fmt.Sprintf("Store should have size %d", i+1))
	}
}
