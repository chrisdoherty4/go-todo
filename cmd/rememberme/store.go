package main

import (
	"github.com/chrisdoherty4/rememberme/pkg/todo"
	"github.com/chrisdoherty4/rememberme/pkg/todo/repo"
)

var store = repo.NewMemoryRepository()

func init() {
	store.Save(todo.NewItem("Walk dog"))
	store.Save(todo.NewItem("Walk cat"))
	store.Save(todo.NewItem("Walk crocodile"))
}
