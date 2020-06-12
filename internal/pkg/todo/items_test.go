package todo_test

import (
	"testing"

	"github.com/chrisdoherty4/rememberme/internal/pkg/todo"
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
