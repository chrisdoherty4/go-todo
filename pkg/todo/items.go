package todo

import "sync"

// Item represents a single todo item.
type Item struct {
	title       string
	description string
	complete    bool
	mutex       sync.Mutex
}

// NewItem creates a new Item instance.
func NewItem(title string) *Item {
	return &Item{
		title: title,
	}
}

// Complete retrieves the completion status of the todo item.
func (t *Item) Complete() bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return t.complete
}

// MarkComplete marks the todo item as complete.
func (t *Item) MarkComplete() {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.complete = true
}

// Title retrieves the todo title
func (t *Item) Title() string {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return t.title
}

// Description retrieves the todo description
func (t *Item) Description() string {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return t.description
}

// SetTitle sets the todo item title.
func (t *Item) SetTitle(title string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.title = title
}

// SetDescription sets the todo item description.
func (t *Item) SetDescription(description string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.description = description
}
