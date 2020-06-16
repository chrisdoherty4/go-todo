package router

import "net/http"

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


