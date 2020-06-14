package todo

import "time"

type item struct {
	CreatedAt   time.Time `json:"createdAt"`
	LastUpdated time.Time `json:"lastUpdated"`
}

// Item represents a single todo item.
type Item struct {
	item
	Title       string `json:"title"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}

// MarkComplete marks the todo item as complete.
func (t *Item) MarkComplete() {
	t.Complete = true
}

// Clone clones the instance it's called on.
func (t Item) Clone() *Item {
	return &Item{
		item: item{
			CreatedAt:   t.item.CreatedAt,
			LastUpdated: t.item.LastUpdated,
		},
		Title:       t.Title,
		Description: t.Description,
		Complete:    t.Complete,
	}
}

// Touch updates the last updated data for the item instance.
func (t *Item) Touch() {
	t.item.LastUpdated = time.Now()
}

// NewItem creates a new item.
// The new item will be initialized with the @title argument and all
// other values will have their 0 initialization.
func NewItem(title string) *Item {
	return &Item{
		item: item{
			CreatedAt:   time.Now(),
			LastUpdated: time.Now(),
		},
		Title: title,
	}
}

// Clone clones an existing todo item into a new memory space.
func Clone(src *Item) *Item {
	return src.Clone()
}
