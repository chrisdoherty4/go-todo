package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/chrisdoherty4/rememberme/pkg/router"
	"github.com/chrisdoherty4/rememberme/pkg/router/route"
)

var addr = ":8080"

func main() {
	log.Printf("Starting server on %v", strings.Split(addr, ":")[1])

	r := router.NewRouter()

	r.Handle(
		router.NewInlineRouteHandler(
			route.Get("/"),
			func(w http.ResponseWriter, _ *http.Request) {
				w.Write([]byte("Inline handler"))
			},
		),
	)

	http.ListenAndServe(addr, r)
}
