package main

import (
	"net/http"

	"github.com/chrisdoherty4/rememberme/pkg/mux"
)

func configureHandlers(r *mux.Router) {
	r.Group("/items", func(rg mux.RouteFactory) {

		rg.Get("/", mux.NewInlineHandler(
			func(w http.ResponseWriter, _ *http.Request, _ *mux.RouteMatch) {
				itemController.List(w)
			},
		))

		rg.Get("/([A-Za-z-]+)", mux.NewInlineHandler(
			func(w http.ResponseWriter, _ *http.Request, rm *mux.RouteMatch) {
				itemController.Show(w, rm)
			},
		))

	})
}
