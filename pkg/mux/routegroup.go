package mux

import (
	"fmt"
	"regexp"
	"strings"
)

// RouteGroup is a Route factory that applies shared attributes to Routes created.
type RouteGroup struct {
	router     *Router
	pathPrefix string
}

// PathPrefix retrieves the path prefix applied to all routes created with
// the RouteGroup instance.
func (g RouteGroup) PathPrefix() string {
	return g.pathPrefix
}

// SetPathPrefix sets the path prefix that is applied to all routes created with
// the RouteGroup instance.
func (g *RouteGroup) SetPathPrefix(prefix string) {
	g.pathPrefix = fmt.Sprintf("/%v", strings.Trim(prefix, "/"))
}

// NewRoute creates a new Route instance using the group attributes.
func (g *RouteGroup) newRoute(method, path string, handler HandlerFunc) *Route {
	path = fmt.Sprintf("%v/%v", g.pathPrefix, path)

	multiSlashRegex := regexp.MustCompile("/{2,}")
	path = strings.TrimRight(multiSlashRegex.ReplaceAllString(path, "/"), "/")

	return g.router.newRoute(method, path, handler)
}

// Clone clones the RouteGroup instance it's called on.
func (g RouteGroup) Clone() *RouteGroup {
	return &g
}

// Get creates a new MethodGet Route using the group attributes.
func (g *RouteGroup) Get(path string, handler HandlerFunc) *Route {
	return g.newRoute(MethodGet, path, handler)
}

// Post creates a new MethodPost Route using the group attributes
func (g *RouteGroup) Post(path string, handler HandlerFunc) *Route {
	return g.newRoute(MethodPost, path, handler)
}

// Put creates a new MethodPut Route using the group attributes
func (g *RouteGroup) Put(path string, handler HandlerFunc) *Route {
	return g.newRoute(MethodPut, path, handler)
}

// Delete creates a new MethodDelete Route using the group attributes
func (g *RouteGroup) Delete(path string, handler HandlerFunc) *Route {
	return g.newRoute(MethodDelete, path, handler)
}

// Group creates a clone of this RouteGroup instance appending the path
// argument to it's path prefix. The clone is then injected into the
// RouteGroupFunc.
func (g *RouteGroup) Group(path string, f RouteFactoryFunc) RouteFactory {
	routeGroup := g.Clone()
	routeGroup.SetPathPrefix(fmt.Sprintf("%v/%v", routeGroup.PathPrefix(), path))

	f(routeGroup)

	return routeGroup
}

// NewRouteGroup creates a new Group instance.
func NewRouteGroup(path string, r *Router) *RouteGroup {
	rg := &RouteGroup{
		router: r,
	}

	rg.SetPathPrefix(path)

	return rg
}
