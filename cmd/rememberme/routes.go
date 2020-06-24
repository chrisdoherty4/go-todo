package main

import (
	"github.com/chrisdoherty4/rememberme/pkg/mux"
)

func configureHandlers(r *mux.Router) {
	r.Group("/items", func(rg mux.RouteFactory) {

		rg.Get("/", mux.NewInlineHandler(itemController.List))

		rg.Get("/([A-Za-z-]+)", mux.NewInlineHandler(itemController.Show))

	})
}
