package main

import (
	"github.com/chrisdoherty4/rememberme/pkg/mux"
)

var itemController = NewItemController(store)

func configureHandlers(r mux.RouteFactory) {

	r.Group("/api/v1", func(rg mux.RouteFactory) {

		rg.Group("/items", func(rg mux.RouteFactory) {

			rg.Get("/", itemController.List)

			rg.Group("/([A-Za-z-]+)", func(rg mux.RouteFactory) {

				rg.Get("/", itemController.Show)

				rg.Post("/", itemController.Save)

				rg.Delete("/", itemController.Delete)

			})

		})

	})
}
