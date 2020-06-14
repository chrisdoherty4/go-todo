/*
Package route contains helper functions for creating router.Route instances
*/
package route

import "github.com/chrisdoherty4/rememberme/pkg/router"

// Get creates a new GET method route.
func Get(path string) *router.Route {
	return router.NewRoute(router.MethodGet, path)
}

// Post creates a new POST method route.
func Post(path string) *router.Route {
	return router.NewRoute(router.MethodPost, path)
}

// Put creates a new PUT method route.
func Put(path string) *router.Route {
	return router.NewRoute(router.MethodPut, path)
}

// Delete creates a new DELETE method route.
func Delete(path string) *router.Route {
	return router.NewRoute(router.MethodDelete, path)
}
