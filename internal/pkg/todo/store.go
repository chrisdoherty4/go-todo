package todo

// Store is an interface for storing todo list items.
type Store interface {
	// Add adds an item to the store
	Add(item *Item)

	// Delete removes an item from the store
	Delete(title string) *Item

	// Get retrieves a todo list item from the store
	Get(title string) *Item

	// GetAll retrieves all items in the store.
	GetAll() []*Item

	// MarkComplete marks an item in the list complete.
	MarkComplete(title string)

	// Size retrieves the total number of items in the store.
	Size() int
}
