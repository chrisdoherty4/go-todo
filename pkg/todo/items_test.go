package todo_test

import (
	"fmt"
	"testing"

	"github.com/chrisdoherty4/go-todo/pkg/todo"
	"github.com/stretchr/testify/assert"
)

func TestItemCreation(t *testing.T) {
	title := "Grocery shopping"
	description := `
		- Apples
		- Bananas
		- Pears
	`

	item := todo.NewItem(title)
	assert.Equal(t, item.Title(), title, "Title should be equal")

	item.SetDescription(description)
	assert.Equal(t, item.Description(), description, "Description should be equal")

	assert.False(t, item.Complete(), "Item should not be complete")
	item.MarkComplete()
	assert.True(t, item.Complete(), "Item should be complete")
}

func TestListCreation(t *testing.T) {
	list := todo.NewList()
	assert.Zero(t, list.Size(), "List should be initialized with 0 items")

	list.Add(todo.NewItem("Walk dog"))
	assert.Equal(t, list.Size(), 1, "List should have 1 item in")
}

func TestAddMultipleItems(t *testing.T) {
	items := []struct {
		title       string
		description string
	}{
		{
			"Grocery shopping",
			`
				- Apples
				- Bananas
				- Pears
			`,
		},
		{"Clean car", "Clean car"},
		{"Walk dog", ""},
	}

	list := todo.NewList()

	for i, data := range items {
		item := todo.NewItem(data.title)
		item.SetDescription(data.description)

		list.Add(item)
		assert.Equal(t, list.Size(), i+1, fmt.Sprintf("List should have size %d", i+1))
	}
}
