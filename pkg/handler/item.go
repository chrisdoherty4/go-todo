package handler

import (
	"encoding/json"
	"net/http"

	"github.com/chrisdoherty4/rememberme/pkg/mux"
	"github.com/chrisdoherty4/rememberme/pkg/todo"
)

// ItemController handles CRUD interactions with a todo repository.
type ItemController struct {
	repo todo.Repository
}

// List lists all items in the todo repository.
func (ic ItemController) List(w http.ResponseWriter, r *http.Request, _ *mux.RouteMatch) {
	items, err := json.MarshalIndent(ic.repo.GetAll(), "", "  ")

	if err != nil {
		http.Error(w, "Could not marshal todo items as json", 500)
	}

	w.Write([]byte(items))
}

// NewItemController creates a new ItemController instance.
func NewItemController(repo todo.Repository) *ItemController {
	return &ItemController{
		repo: repo,
	}
}
