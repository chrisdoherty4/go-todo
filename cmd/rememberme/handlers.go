package main

import (
	"encoding/json"
	"net/http"

	"github.com/chrisdoherty4/rememberme/pkg/router"
	"github.com/chrisdoherty4/rememberme/pkg/router/route"
	"github.com/chrisdoherty4/rememberme/pkg/todo"
)

type listItemsHandler struct {
	repo todo.Repository
}

func (t listItemsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	items, err := json.MarshalIndent(t.repo.GetAll(), "", "  ")

	if err != nil {
		http.Error(w, "Could not marshal todo items as json", 500)
	}

	w.Write([]byte(items))
}

func newListItemsHandler(repo todo.Repository) *listItemsHandler {
	return &listItemsHandler{
		repo: repo,
	}
}

func configureHandlers(r *router.Router) {
	root := route.NewGroup()
	root.SetPathPrefix("/api/v1")

	r.Handle(router.NewRouteHandler(
		root.Get("/items"),
		newListItemsHandler(store),
	))
}
