package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/chrisdoherty4/rememberme/pkg/router"
	"github.com/chrisdoherty4/rememberme/pkg/router/route"
	"github.com/chrisdoherty4/rememberme/pkg/todo"
	"github.com/chrisdoherty4/rememberme/pkg/todo/repo"
)

var host = ""
var addr = ":8080"

var store = repo.NewMemoryRepository()

func init() {
	store.Save(todo.NewItem("Walk dog"))
	store.Save(todo.NewItem("Walk cat"))
	store.Save(todo.NewItem("Walk crocodile"))
}

func main() {
	log.Printf("Server listening at %v", strings.Split(addr, ":")[1])

	r := router.NewRouter()

	rg := route.NewGroup()
	rg.SetPathPrefix("/api/v1")

	r.Handle(router.NewRouteHandler(
		rg.Get("/items"),
		newListItemsHandler(store),
	))

	log.Fatal(http.ListenAndServe(addr, r))
}
