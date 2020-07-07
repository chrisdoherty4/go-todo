package mux

import (
	"log"
	"net/http"
)

// Router dispatches requests to handlers based on route matches.
type Router struct {
	routes []*Route
}

// ServeHTTP handles requests received by the Http server.
// See https://golang.org/pkg/net/http/#Handler
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	log.Printf("[%v] %v %v", request.RemoteAddr, request.Method, request.URL.Path)

	route := r.FindRoute(request)

	if route != nil {
		route.Handle(w, request)
		return
	}

	// Make this customizable
	http.NotFound(w, request)
}

func (r *Router) newRoute(method, path string, handler HandlerFunc) *Route {
	route := NewRoute(method, path, NewInlineHandler(handler))
	r.routes = append(r.routes, route)

	return route
}

// Get creates a new MethodGet Route using the group attributes.
func (r *Router) Get(path string, handler HandlerFunc) *Route {
	return r.newRoute(MethodGet, path, handler)
}

// Post creates a new MethodPost Route using the group attributes
func (r *Router) Post(path string, handler HandlerFunc) *Route {
	return r.newRoute(MethodPost, path, handler)
}

// Put creates a new MethodPut Route using the group attributes
func (r *Router) Put(path string, handler HandlerFunc) *Route {
	return r.newRoute(MethodPut, path, handler)
}

// Delete creates a new MethodDelete Route using the group attributes
func (r *Router) Delete(path string, handler HandlerFunc) *Route {
	return r.newRoute(MethodDelete, path, handler)
}

// Count retrieves the total number of routes registered with the Router.
func (r Router) Count() int {
	return len(r.routes)
}

// FindRoute retrieves a route from the router that matches a request.
func (r Router) FindRoute(request *http.Request) *Route {
	for _, route := range r.routes {
		if route.Match(request) {
			return route
		}
	}

	return nil
}

// Group creates a new RouteGroup that will register routes with this Route
// instance.
func (r *Router) Group(path string, f RouteFactoryFunc) RouteFactory {
	routeGroup := NewRouteGroup(path, r)
	routeGroup.SetPathPrefix(path)

	f(routeGroup)

	return routeGroup
}

// NewRouter creates a new Router instance.
func NewRouter() *Router {
	return &Router{}
}
