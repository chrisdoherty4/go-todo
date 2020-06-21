package handler

import (
	"encoding/json"
	"net/http"

	"github.com/chrisdoherty4/rememberme/pkg/todo"
)

// ListItemsHandler is a handler for listing all items in a todo repository.
type ListItemsHandler struct {
	repo todo.Repository
}

// Handle lists all items in the todo repository.
func (t ListItemsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	items, err := json.MarshalIndent(t.repo.GetAll(), "", "  ")

	if err != nil {
		http.Error(w, "Could not marshal todo items as json", 500)
	}

	w.Write([]byte(items))
}

// NewListItemsHandler creates a new ListItemsHandler instance.
func NewListItemsHandler(repo todo.Repository) *ListItemsHandler {
	return &ListItemsHandler{
		repo: repo,
	}
}
