package router

// RouteFactory creates routes.
type RouteFactory interface {
	Get(string) *router.Route

	Post(string) *router.Route

	Put(string) *router.Route

	Delete(string) *router.Route
}
