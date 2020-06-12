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

// List is a todo list containing items.
type List struct {
	items []*Item
	mutex sync.Mutex
}

// NewList creates a new List instance.
func NewList() *List {
	return &List{}
}

// Add adds an item to the list
func (t *List) Add(item *Item) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.items = append(t.items, item)
}

// Remove removes an item from the list
func (t *List) Remove(title string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	for i, item := range t.items {
		if item.title == title {
			t.items = append(t.items[i:], t.items[i+1:]...)
			return
		}
	}
}

// MarkComplete marks an item in the list complete.
// If the item does not exist in the list the call is a noop.
func (t *List) MarkComplete(title string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	for _, item := range t.items {
		if item.title == title {
			item.MarkComplete()
			return
		}
	}
}

// Size retrieves the total number of items in the List.
func (t *List) Size() int {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return len(t.items)
}
