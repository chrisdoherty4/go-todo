package store

import (
	"sync"

	"github.com/chrisdoherty4/go-todo/pkg/todo"
)

// MemoryStore is an in memory storage structure for todo list items.
type MemoryStore struct {
	items []*todo.Item
	mutex sync.Mutex
}

// NewMemoryStore creates a MemoryStore instance.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

// Add adds a todo item to the MemoryStore.
func (t *MemoryStore) Add(item *todo.Item) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.items = append(t.items, item)
}

// Delete removes a todo item from the MemoryStore.
//
// If the item is not in the store a nil pointer is returned.
func (t *MemoryStore) Delete(title string) *todo.Item {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	for i, item := range t.items {
		if item.Title() == title {
			t.items = append(t.items[i:], t.items[i+1:]...)
			return item
		}
	}

	return nil
}

// Get retrieves a todo MemoryStore todo. from the MemoryStore.
//
// If the item is not in the store a nil pointer is returned.
func (t *MemoryStore) Get(title string) *todo.Item {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	for _, item := range t.items {
		if item.Title() == title {
			return item
		}
	}

	return nil
}

// GetAll retrieves all todo.s in the MemoryStore.
func (t *MemoryStore) GetAll() []*todo.Item {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// Create a copy of the data and return a slice to that so as to
	// preserve the integrity of the data managed by the MemoryStore
	// instance.
	c := make([]*todo.Item, len(t.items), cap(t.items))
	copy(c, t.items)
	return c
}

// MarkComplete marks an todo. in the MemoryStore complete.
// If the todo. does not exist in the MemoryStore the call is a noop.
func (t *MemoryStore) MarkComplete(title string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	for _, item := range t.items {
		if item.Title() == title {
			item.MarkComplete()
			return
		}
	}
}

// Size retrieves the total number of todo.s in the MemoryStore.
func (t *MemoryStore) Size() int {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return len(t.items)
}
