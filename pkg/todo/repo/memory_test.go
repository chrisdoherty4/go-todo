package repo_test

import (
	"fmt"
	"testing"

	"github.com/chrisdoherty4/rememberme/pkg/todo"
	"github.com/chrisdoherty4/rememberme/pkg/todo/repo"
	"github.com/stretchr/testify/assert"
)

var title = "Walk dog"
var item = todo.NewItem(title)

func TestMemoryRepository(t *testing.T) {
	repo := repo.NewMemoryRepository()
	assert.Zero(t, repo.Size(), "List should be initialized with 0 items")

	repo.Save(todo.NewItem("Walk dog"))
	assert.Equal(t, repo.Size(), 1, "List should have 1 item in")
}

func TestSaveMultipleItems(t *testing.T) {
	items := []struct {
		title string
	}{
		{"Grocery shopping"},
		{"Clean car"},
		{"Walk dog"},
	}

	repo := repo.NewMemoryRepository()

	for i, data := range items {
		item := todo.NewItem(data.title)
		repo.Save(item)
		assert.Equal(t, repo.Size(), i+1, fmt.Sprintf("Repository should have size %d", i+1))
	}
}

func TestMarkItemComplete(t *testing.T) {
	repo := repo.NewMemoryRepository()

	repo.Save(item)
	repo.MarkComplete(title)
	assert.True(t, repo.Get(title).Complete, "Item should have been marked complete")
}

func TestDeletingItems(t *testing.T) {
	repo := repo.NewMemoryRepository()

	repo.Save(item)
	assert.Equal(t, 1, repo.Size(), "Repo should have 1 item in.")
	repo.Delete(title)
	assert.Equal(t, 0, repo.Size(), "Repo should have 0 items in.")
}

func TestGettingItem(t *testing.T) {
	repo := repo.NewMemoryRepository()

	description := "Walk the damn dog"
	item.Description = description
	repo.Save(item)
	retrieved := repo.Get(title)
	assert.Equal(t, description, retrieved.Description, fmt.Sprintf("The description should be '%v'", description))
}
