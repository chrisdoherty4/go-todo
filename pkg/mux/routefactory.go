package mux

// RouteFactoryFunc is an interface to bundle related routes under shared
// route settings.
type RouteFactoryFunc func(RouteFactory)

// RouteFactory creates routes.
type RouteFactory interface {
	Get(string, HandlerFunc) *Route

	Post(string, HandlerFunc) *Route

	Put(string, HandlerFunc) *Route

	Delete(string, HandlerFunc) *Route

	Group(string, RouteFactoryFunc) RouteFactory
}
