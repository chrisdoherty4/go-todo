package repo

import (
	"sync"

	"github.com/chrisdoherty4/rememberme/pkg/todo"
)

// MemoryRepository is an in memory storage structure for todo list items.
type MemoryRepository struct {
	items []*todo.Item
	mutex sync.Mutex
}

// Save adds a todo item to the MemoryRepository.
func (t *MemoryRepository) Save(item *todo.Item) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// item.Touch()

	t.items = append(t.items, item)
}

// Delete removes a todo item from the MemoryRepository.
//
// If the item is not in the repo a nil pointer is returned.
func (t *MemoryRepository) Delete(title string) *todo.Item {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	for i, item := range t.items {
		if item.Title == title {
			t.items[i] = t.items[len(t.items)-1]
			t.items = t.items[:len(t.items)-1]
			return item
		}
	}

	return nil
}

// Get retrieves a todo MemoryRepository todo. from the MemoryRepository.
//
// If the item is not in the repo a nil pointer is returned.
func (t *MemoryRepository) Get(title string) *todo.Item {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	for _, item := range t.items {
		if item.Title == title {
			// Return a copy of the item so the caller doesn't try to manipulate
			// the MemoryStore's copy.
			return todo.Clone(item)
		}
	}

	return nil
}

// GetAll retrieves all items in the repository.
func (t *MemoryRepository) GetAll() []*todo.Item {
	// TODO: Think about this more... the algorithm sucks as it's O(N^2) space
	// and time. Perhaps the GetAll interface needs changing to a stream based
	// approach or paging.
	items := make([]*todo.Item, len(t.items))
	copy(items, t.items)

	return items
}

// MarkComplete marks an todo. in the MemoryRepository complete.
// If the todo. does not exist in the MemoryRepository the call is a noop.
//
// TODO: Remove this, it's not a function of a repository.
func (t *MemoryRepository) MarkComplete(title string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	for _, item := range t.items {
		if item.Title == title {
			item.MarkComplete()
			// item.Touch()
			return
		}
	}
}

// Size retrieves the total number of todo.s in the MemoryRepository.
func (t *MemoryRepository) Size() int {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return len(t.items)
}

// NewMemoryRepository creates a MemoryRepository instance.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{}
}
