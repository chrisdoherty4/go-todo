package todo

// Item represents a single todo item.
type Item struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
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

// Clone clones the instance it's called on.
func (t *Item) Clone() *Item {
	return &Item{
		Title:       t.Title,
		Description: t.Description,
		Complete:    t.Complete,
	}
}

// Clone clones an existing todo item into a new memory space.
func Clone(src *Item) *Item {
	return src.Clone()
}
