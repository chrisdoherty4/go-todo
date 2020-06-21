package mux_test

import (
	"net/http"

	"github.com/chrisdoherty4/rememberme/pkg/mux"
)

func ExampleRouter_Get() {
	r := mux.NewRouter()

	r.Get("/", mux.NewInlineHandler(
		func(_ http.ResponseWriter, _ *http.Request) {},
	))
}

func ExampleRouter_Post() {
	r := mux.NewRouter()

	r.Post("/", mux.NewInlineHandler(
		func(_ http.ResponseWriter, _ *http.Request) {},
	))
}

func ExampleRouter_Put() {
	r := mux.NewRouter()

	r.Put("/", mux.NewInlineHandler(
		func(_ http.ResponseWriter, _ *http.Request) {},
	))
}

func ExampleRouter_Delete() {
	r := mux.NewRouter()

	r.Delete("/", mux.NewInlineHandler(
		func(_ http.ResponseWriter, _ *http.Request) {},
	))
}

func ExampleRouter_Group() {
	r := mux.NewRouter()

	r.Group("/example", func(g *mux.RouteGroup) {
		g.Get("/", mux.NewInlineHandler(
			func(_ http.ResponseWriter, _ *http.Request) {},
		))
	})
}
