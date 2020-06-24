package main

import (
	"github.com/chrisdoherty4/rememberme/cmd/rememberme/handler"
	"github.com/chrisdoherty4/rememberme/pkg/mux"
)

var itemController = handler.NewItemController(store)

func configureHandlers(r mux.RouteFactory) {

	r.Group("/api/v1", func(rg mux.RouteFactory) {

		rg.Group("/items", func(rg mux.RouteFactory) {

			rg.Get("/", mux.NewInlineHandler(itemController.List))

			rg.Group("/([A-Za-z-]+)", func(rg mux.RouteFactory) {

				rg.Get("/", mux.NewInlineHandler(itemController.Show))

				rg.Post("/", mux.NewInlineHandler(itemController.Save))

			})

		})

	})
}
