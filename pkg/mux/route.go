package mux

import (
	"errors"
	"net/http"
	"regexp"
)

// Route defines
type Route struct {
	method       string
	path         string
	compiledPath *regexp.Regexp
	handler      Handler
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

	if !r.compiledPath.MatchString(request.URL.Path) {
		return false
	}

	return true
}

// Handle is used by a Router to route a request to the Route's handler.
// Routes should have been tested against the request using the Route.Match()
// method. Failure to do so could result in incorrect token extractions
// of the http.Request path.
func (r *Route) Handle(writer http.ResponseWriter, request *http.Request) {
	// Use [1:] to remove the matching string which is index 0.
	matches := r.compiledPath.FindStringSubmatch(request.URL.Path)[1:]

	r.handler.Handle(writer, request, NewRouteMatch(r, request, matches))
}

// NewRoute creates a new Route instance.
func NewRoute(method, path string, h Handler) *Route {
	return &Route{
		method:       method,
		path:         path,
		compiledPath: regexp.MustCompile(path),
		handler:      h,
	}
}

// RouteMatch wraps a Route and all tokens extracted from the Path.
type RouteMatch struct {
	Route   *Route
	Request *http.Request
	matches []string
}

// Var retrieves a matched variable from the request path at the specified
// index.
func (rm RouteMatch) Var(idx int) (string, error) {
	if idx >= len(rm.matches) {
		return "", errors.New("Index out of bounds")
	}

	return rm.matches[idx], nil
}

// Count returns the total number of matches made against the
func (rm RouteMatch) Count() int {
	return len(rm.matches)
}

// NewRouteMatch creates a new RouteMatch instance.
func NewRouteMatch(r *Route, request *http.Request, m []string) *RouteMatch {
	return &RouteMatch{
		Route:   r,
		Request: request,
		matches: m,
	}
}
