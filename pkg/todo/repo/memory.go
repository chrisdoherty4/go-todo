package repo

import (
	"fmt"
	"sync"

	"github.com/chrisdoherty4/rememberme/pkg/todo"
)

// MemoryRepository is an in memory storage structure for todo list items.
type MemoryRepository struct {
	items []*todo.Item
	mutex sync.Mutex
}

// Save adds a todo item to the MemoryRepository.
func (t *MemoryRepository) Save(source *todo.Item) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	target := t.find(source.Title())

	if target != nil {
		update(target, source)
		return
	}

	t.items = append(t.items, source.Clone())
}

// Delete removes a todo item from the MemoryRepository.
//
// If the item is not in the repo a nil pointer is returned.
func (t *MemoryRepository) Delete(title string) (todo.Item, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	for i, item := range t.items {
		if item.Title() == title {
			t.items[i] = t.items[len(t.items)-1]
			t.items = t.items[:len(t.items)-1]
			return *item, nil
		}
	}

	var item todo.Item
	return item, fmt.Errorf("Item not found (%v)", title)
}

// Get retrieves a todo MemoryRepository todo. from the MemoryRepository.
//
// If the item is not in the repo a nil pointer is returned.
func (t *MemoryRepository) Get(title string) (todo.Item, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	item := t.find(title)

	if item == nil {
		var item todo.Item
		return item, fmt.Errorf("Item not found (%v)", title)
	}

	return *item, nil
}

// GetAll will concurrently dispatch all items to a channel.
func (t *MemoryRepository) GetAll() <-chan todo.Item {
	output := make(chan todo.Item)

	go func() {
		defer close(output)
		for _, item := range t.items {
			output <- *item
		}
	}()

	return output
}

// Size retrieves the total number of todo.s in the MemoryRepository.
func (t *MemoryRepository) Size() int {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return len(t.items)
}

func (t *MemoryRepository) find(title string) *todo.Item {
	for _, item := range t.items {
		if item.Title() == title {
			// Return a copy of the item so the caller doesn't try to manipulate
			// the MemoryStore's copy.
			return item
		}
	}

	return nil
}

// NewMemoryRepository creates a MemoryRepository instance.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{}
}

func update(target, source *todo.Item) {
	target.SetDescription(source.Description())

	if source.Complete() {
		target.MarkComplete()
	}
}
