package todo

// item masks fields in the todo.Item struct. We're doing this to ensure the
// json package can still identify, through reflection, fields that need
// to be included as part of the json.Marshal() call. However, accessing
// the embedded fields directly is prohibited favoring use of the
// accessors and mutators instead.
//
// E.g.
//		item := NewItem("Walk dog")
//		item.Title = "Walk cat" // prohibited even though it's possible.
//
// instead
//
//		item := NewItem("Walk dog")
//		item.SetTitle("Walk cat")
//
// Note: this implementation actually prevents accessing the item fields
// because the Item.Title() and similar symbols take presedence.
type item struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}

// Item represents a single todo item.
type Item struct {
	item
}

// MarkComplete marks the todo item as complete.
func (t *Item) MarkComplete() {
	t.item.Complete = true
}

// Clone clones the instance it's called on.
func (t Item) Clone() *Item {
	return &Item{
		item: item{
			Title:       t.Title(),
			Description: t.Description(),
			Complete:    t.Complete(),
		},
	}
}

// Title retrieves the Item title.
func (t Item) Title() string {
	return t.item.Title
}

// SetTitle sets the title of the Item.
func (t *Item) SetTitle(title string) {
	t.item.Title = title
}

// Description retrieves the Item description.
func (t Item) Description() string {
	return t.item.Description
}

// SetDescription sets the description of the Item.
func (t *Item) SetDescription(desc string) {
	t.item.Description = desc
}

// Complete determines whether the item has been marked complete.
func (t Item) Complete() bool {
	return t.item.Complete
}

// Clone clones an existing todo item into a new memory space.
func Clone(src *Item) *Item {
	return src.Clone()
}

// Touch updates the last updated data for the item instance.
// func (t *Item) Touch() {
// 	t.item.LastUpdated = time.Now()
// }

// NewItem creates a new item.
// The new item will be initialized with the @title argument and all
// other values will have their 0 initialization.
func NewItem(title string) *Item {
	return &Item{
		item: item{
			Title: title,
		},
	}
}
