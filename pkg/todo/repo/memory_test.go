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

func TestCreatingAMemoryRepository(t *testing.T) {
	repo := repo.NewMemoryRepository()
	assert.Zero(t, repo.Size(), "List should be initialized with 0 items")

	repo.Save(todo.NewItem("Walk dog"))
	assert.Equal(t, 1, repo.Size(), "List should have 1 item in")
}

func TestSavingAnItem(t *testing.T) {
	repo := repo.NewMemoryRepository()

	repo.Save(item)
	retrieved, err := repo.Get(title)

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, item.Title(), retrieved.Title(), "Titles should be the same")

	description := "Walk the dog around the block"
	retrieved.SetDescription(description)
	repo.Save(&retrieved)
	retrieved, err = repo.Get(title)

	if err != nil {
		t.Fail()
	}

	assert.Equal(
		t,
		description,
		retrieved.Description(),
		"The descriptions should be the same",
	)
}

func TestSavingMultipleItems(t *testing.T) {
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
		assert.Equal(
			t,
			i+1,
			repo.Size(),
			fmt.Sprintf("Repository should have size %d", i+1),
		)
	}
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

	repo.Save(item)

	retrieved, err := repo.Get(title)

	if err != nil {
		t.Fail()
	}

	assert.Equal(
		t,
		title,
		retrieved.Title(),
		fmt.Sprintf("The title should be '%v'", title),
	)
}
