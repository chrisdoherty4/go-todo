package router

import (
	"fmt"
	"net/http"
	"strings"
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

// NewRoute creates a new Route instance using the group attributes.
func (r *Route) NewRoute(method, path string) *router.Route {
	if g.pathPrefix != "" {
		path = fmt.Sprintf("%v/%v", g.pathPrefix, strings.Trim(path, "/"))
	}

	return NewRoute(method, path)
}

// Get creates a new router.MethodGet Route using the group attributes.
func (r Route) Get(path string) *Route {
	return g.NewRoute(MethodGet, path)
}

// Post creates a new router.MethodPost Route using the group attributes
func (r Route) Post(path string) *Route {
	return g.NewRoute(MethodPost, path)
}

// Put creates a new router.MethodPut Route using the group attributes
func (r Route) Put(path string) *Route {
	return g.NewRoute(MethodPut, path)
}

// Delete creates a new router.MethodDelete Route using the group attributes
func (r Route) Delete(path string) *Route {
	return r.NewRoute(MethodDelete, path)
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

func (r *Route) Route(writer http.ResponseWriter, request *http.Request) {
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
