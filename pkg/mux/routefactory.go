package mux

// RouteFactoryFunc is an interface to bundle related routes under shared
// route settings.
type RouteFactoryFunc func(RouteFactory)

// RouteFactory creates routes.
type RouteFactory interface {
	Get(string, Handler) *Route

	Post(string, Handler) *Route

	Put(string, Handler) *Route

	Delete(string, Handler) *Route

	Group(string, RouteFactoryFunc) RouteFactory
}
