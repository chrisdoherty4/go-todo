package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chrisdoherty4/rememberme/pkg/handler"
	"github.com/chrisdoherty4/rememberme/pkg/mux"
)

var host = ""
var port = 8080
var addr = fmt.Sprintf("%v:%v", host, port)

func main() {
	log.Printf("Server listening at %v", addr)

	r := mux.NewRouter()
	configureHandlers(r)

	log.Fatal(http.ListenAndServe(addr, r))
}

func configureHandlers(r *mux.Router) {
	r.Group("/items", func(rg *mux.RouteGroup) {

		rg.Get("/", handler.NewListItemsHandler(store))

	})
}
