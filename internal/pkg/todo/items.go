package todo

// Item represents a single todo item.
type Item struct {
	Title       string
	Description string
	Complete    bool
}

// MarkComplete marks the todo item as complete.
func (t *Item) MarkComplete() {
	t.Complete = true
}

// NewItem creates a new item.
// The new item will be initialized with the @title argument and all
// other values will have their 0 initialization.
func NewItem(title string) *Item {
	return &Item{
		Title: title,
	}
}

// Clone clones an existing todo item into a new memory space.
func Clone(src *Item) *Item {
	return &Item{
		Title:       src.Title,
		Description: src.Description,
		Complete:    src.Complete,
	}
}
