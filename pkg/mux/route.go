package mux

import (
	"net/http"
)

// Route defines
type Route struct {
	method  string
	path    string
	handler Handler
}

// Method retrieves the http method associated with this route.
func (r *Route) Method() string {
	return r.method
}

// Path retrieves the URL path this route matches against.
func (r *Route) Path() string {
	return r.path
}

// Handler returns the handler associated with this route.
func (r *Route) Handler() Handler {
	return r.handler
}

// Equal compares this Route instance against another to determine if they are
// equal. Route instances are equal if the method and path of each route
// are equal.
func (r Route) Equal(e *Route) bool {
	return r.method == e.path && r.path == e.path
}

// Match is used by the Router to match requests to Routes.
func (r Route) Match(request *http.Request) bool {
	// TODO: Create more sophisticated algorithm for handling path tokens
	if r.method != request.Method {
		return false
	}

	if r.path != request.URL.Path {
		return false
	}

	return true
}

// Handle is used by a Router to route a request to the Route's handler.
func (r *Route) Handle(writer http.ResponseWriter, request *http.Request) {
	r.handler.Handle(writer, request)
}

// NewRoute creates a new Route instance.
func NewRoute(method, path string, h Handler) *Route {
	return &Route{
		method:  method,
		path:    path,
		handler: h,
	}
}
