package mux_test

import (
	"net/http"

	"github.com/chrisdoherty4/rememberme/pkg/mux"
)

func ExampleRouter_Get() {
	r := mux.NewRouter()

	r.Get("/", mux.NewInlineHandler(
		func(_ http.ResponseWriter, _ *http.Request, _ *mux.RouteMatch) {},
	))
}

func ExampleRouter_Post() {
	r := mux.NewRouter()

	r.Post("/", mux.NewInlineHandler(
		func(_ http.ResponseWriter, _ *http.Request, _ *mux.RouteMatch) {},
	))
}

func ExampleRouter_Put() {
	r := mux.NewRouter()

	r.Put("/", mux.NewInlineHandler(
		func(_ http.ResponseWriter, _ *http.Request, _ *mux.RouteMatch) {},
	))
}

func ExampleRouter_Delete() {
	r := mux.NewRouter()

	r.Delete("/", mux.NewInlineHandler(
		func(_ http.ResponseWriter, _ *http.Request, _ *mux.RouteMatch) {},
	))
}

func ExampleRouter_Group() {
	r := mux.NewRouter()

	// Create a new RouteGroup instance and bundle the configurations
	// in the anonymous function. The function adheres to the RouteGroupFunc
	// type.
	r.Group("/example", func(rf mux.RouteFactory) {

		// Define routes that are prefixed with the route group /example path.
		rf.Get("/", mux.NewInlineHandler(
			func(_ http.ResponseWriter, _ *http.Request, _ *mux.RouteMatch) {},
		))

	})
}
