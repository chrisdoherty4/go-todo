package todo_test

import (
	"testing"

	"github.com/chrisdoherty4/rememberme/internal/pkg/todo"
	"github.com/stretchr/testify/assert"
)

func TestClone(t *testing.T) {
	item := &todo.Item{
		Title: "Walk dog",
	}

	c := todo.Clone(item)

	assert.True(t, item != c, "Items should have different memory spaces.")
}

func TestMarkComplete(t *testing.T) {
	item := &todo.Item{
		Title: "Walk dog",
	}

	assert.False(t, item.Complete, "Item should be incomplete")
	item.MarkComplete()
	assert.True(t, item.Complete, "Item should be complete")
}
