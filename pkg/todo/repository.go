package todo

// Repository is an interface for storing todo list items.
type Repository interface {
	// Save adds an item to the repository or updates an existing item in the
	// repository
	Save(item *Item)

	// Delete removes an item from the repository
	// The returned Item is the item just deleted.
	Delete(title string) (Item, error)

	// Get retrieves a todo list item from the repository
	Get(title string) (Item, error)

	// GetAll retrieves all items in the repository.
	GetAll() <-chan Item

	// Size retrieves the total number of items in the repository.
	Size() int
}
