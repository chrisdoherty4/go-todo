package router

import (
	"log"
	"net/http"
)

// RouteHandler binds a Route instance with a Handler.
type RouteHandler struct {
	route   *Route
	handler Handler
}

// Route retrieves the route.
func (t RouteHandler) Route() *Route {
	return t.route
}

// Handler retrieves the Handler instance.
func (t RouteHandler) Handler() Handler {
	return t.handler
}

// Match determines if a request matches this RouteHandler instance.
func (t RouteHandler) Match(r *http.Request) bool {
	return t.route.Match(r)
}

// Handle handles a request.
// The RouteHandler should be matched to a request first by checking
// RouteHandler.Match(...)
func (t RouteHandler) Handle(w http.ResponseWriter, r *http.Request) {
	t.handler.Handle(w, r)
}

// NewInlineRouteHandler creates a new RouteHandler instance.
func NewInlineRouteHandler(r *Route, h InlineHandler) *RouteHandler {
	return &RouteHandler{
		route: r,
		handler: &inlineHandlerWrapper{
			handler: h,
		},
	}
}

// NewRouteHandler creates a new RouteHandler instance.
func NewRouteHandler(r *Route, h Handler) *RouteHandler {
	return &RouteHandler{
		route:   r,
		handler: h,
	}
}

// Router dispatches requests to handlers based on route matches.
type Router struct {
	handlers []*RouteHandler
}

// Handle a route using a predefined handler.
func (t *Router) Handle(rh *RouteHandler) {
	// TODO: Add protection for 1-2-many route-to-handlers.
	t.handlers = append(t.handlers, rh)
}

// ServeHTTP handles requests received by the Http server.
// See https://golang.org/pkg/net/http/#Handler
func (t *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%v] %v %v", r.RemoteAddr, r.Method, r.URL.Path)

	for _, handler := range t.handlers {
		if handler.Match(r) {
			handler.Handle(w, r)
			return
		}
	}

	http.NotFound(w, r)
}

// NewRouter creates a new Router instance.
func NewRouter() *Router {
	return &Router{}
}
