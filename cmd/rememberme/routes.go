package main

import (
	"net/http"

	"github.com/chrisdoherty4/rememberme/pkg/mux"
)

func configureHandlers(r *mux.Router) {
	r.Group("/items", func(rg *mux.RouteGroup) {

		rg.Get("/", mux.NewInlineHandler(
			func(w http.ResponseWriter, r *http.Request, rm *mux.RouteMatch) {
				itemController.List(w, r, rm)
			},
		))

	})
}
