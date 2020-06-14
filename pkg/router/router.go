package router

import (
	"log"
	"net/http"
)

const (
	// MethodGet GET
	MethodGet = "GET"

	// MethodPost POST
	MethodPost = "POST"

	// MethodPut PUT
	MethodPut = "PUT"

	// MethodDelete DELETE
	MethodDelete = "DELETE"
)

// Route defines
type Route struct {
	Method string
	Path   string
}

// Match is used by the Router to match requests to Routes.
func (t *Route) Match(r *http.Request) bool {
	// TODO: Create more sophisticated algorithm for handling path tokens
	if r.Method != t.Method {
		return false
	}

	if r.URL.Path != t.Path {
		return false
	}

	return true
}

// NewRoute creates a new Route instance.
func NewRoute(method, path string) *Route {
	return &Route{
		Method: method,
		Path:   path,
	}
}

// GetRoute creates a new GET method route.
func GetRoute(path string) *Route {
	return NewRoute(MethodGet, path)
}

// PostRoute creates a new POST method route.
func PostRoute(path string) *Route {
	return NewRoute(MethodPost, path)
}

// PutRoute creates a new PUT method route.
func PutRoute(path string) *Route {
	return NewRoute(MethodPut, path)
}

// NewDeleteRoute creates a new DELETE method route.
func NewDeleteRoute(path string) *Route {
	return NewRoute(MethodDelete, path)
}

// Handler defines an interface to handle requests.
type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

// InlineHandler is a type for writing an inline function handler.
type InlineHandler func(http.ResponseWriter, *http.Request)

// Router dispatches requests to handlers based on route matches.
type Router struct {
	routes       map[*Route]Handler
	inlineRoutes map[*Route]InlineHandler
}

// Handle a route using a predefined handler.
func (t *Router) Handle(r *Route, h Handler) {
	t.routes[r] = h
}

// HandleInline handles a route using an inline handler.
func (t *Router) HandleInline(r *Route, h InlineHandler) {
	t.inlineRoutes[r] = h
}

// ServeHTTP handles requests received by the Http server.
// See https://golang.org/pkg/net/http/#Handler
func (t *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%v] %v %v", r.RemoteAddr, r.Method, r.URL.Path)

	// TODO: Add protection for inline handlers and full handlers handling
	// the same routes.
	for route, handler := range t.routes {
		if route.Match(r) {
			handler.Handle(w, r)
			return
		}
	}

	for route, handler := range t.inlineRoutes {
		if route.Match(r) {
			handler(w, r)
			return
		}
	}

	http.NotFound(w, r)
}

// NewRouter creates a new Router instance.
func NewRouter() *Router {
	return &Router{
		routes:       make(map[*Route]Handler),
		inlineRoutes: make(map[*Route]InlineHandler),
	}
}
