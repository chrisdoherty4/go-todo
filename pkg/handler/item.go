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
func (ic ItemController) List(w http.ResponseWriter, _ *http.Request, rm *mux.RouteMatch) {
	items, err := json.MarshalIndent(ic.repo.GetAll(), "", "  ")

	if err != nil {
		http.Error(w, "Could not marshal todo items as json", 500)
	}

	w.Write([]byte(items))
}

// Show retrieves a single item from the repository and returns it as json.
func (ic ItemController) Show(w http.ResponseWriter, _ *http.Request, rm *mux.RouteMatch) {
	title, _ := rm.Var(0)

	item, err := json.MarshalIndent(ic.repo.Get(title), "", "  ")

	if err != nil {
		http.Error(w, "Could not marshal todo items as json", 500)
	}

	w.Write([]byte(item))
}

// NewItemController creates a new ItemController instance.
func NewItemController(repo todo.Repository) *ItemController {
	return &ItemController{
		repo: repo,
	}
}
