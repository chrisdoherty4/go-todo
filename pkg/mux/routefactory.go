package mux

// RouteGroupFunc is an interface to bundle related routes under shared
// route settings.
type RouteGroupFunc func(*RouteGroup)

// RouteFactory creates routes.
type RouteFactory interface {
	Get(string, Handler) *Route

	Post(string, Handler) *Route

	Put(string, Handler) *Route

	Delete(string, Handler) *Route

	Group(string, RouteGroupFunc) *RouteGroup
}
