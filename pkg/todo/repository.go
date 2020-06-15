package todo

// Repository is an interface for storing todo list items.
type Repository interface {
	// Save adds an item to the repository
	Save(item *Item)

	// Delete removes an item from the repository
	// The item returned should be the item just deleted.
	Delete(title string) *Item

	// Get retrieves a todo list item from the repository
	Get(title string) *Item

	// GetAll retrieves all items in the repository.
	GetAll() []*Item

	// Size retrieves the total number of items in the repository.
	Size() int
}
