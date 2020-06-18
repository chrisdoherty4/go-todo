package router

import "net/http"

// Route defines
type Route struct {
	Routable
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

// Match is used by the Router to match requests to Routes.
func (r *Route) Match(req *http.Request) bool {
	// TODO: Create more sophisticated algorithm for handling path tokens
	if r.method != req.Method {
		return false
	}

	if r.URL.Path != r.path {
		return false
	}

	return true
}

func (r *Route) Route(w http.ResponseWriter, r *http.Request) {
	r.handler.Handle(w, r)
}

// NewRoute creates a new Route instance.
func NewRoute(method, path string, h Handler) *Route {
	return &Route{
		method:  method,
		path:    path,
		handler: h,
	}
}
