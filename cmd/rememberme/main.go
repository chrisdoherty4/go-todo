package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chrisdoherty4/rememberme/pkg/router"
)

var host = ""
var port = 8080
var addr = fmt.Sprintf("%v:%v", host, port)

func main() {
	log.Printf("Server listening at %v", addr)

	r := router.NewRouter()
	configureHandlers(r)

	log.Fatal(http.ListenAndServe(addr, r))
}

func configureHandlers(r *router.Router) {
	root := route.NewGroup()
	root.SetPathPrefix("/api/v1")

	r.Handle(router.NewRouteHandler(
		root.Get("/items"),
		newListItemsHandler(store),
	))
}
