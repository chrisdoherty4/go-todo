package todo

// Repository is an interface for storing todo list items.
type Repository interface {
	// Save adds an item to the store
	Save(item *Item)

	// Delete removes an item from the store
	// The item returned should be the item just deleted.
	Delete(title string) *Item

	// Get retrieves a todo list item from the store
	Get(title string) *Item

	// MarkComplete marks an item in the list complete.
	MarkComplete(title string)

	// Size retrieves the total number of items in the store.
	Size() int
}
