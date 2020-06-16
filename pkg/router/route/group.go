package route

import (
	"fmt"
	"strings"

	"github.com/chrisdoherty4/rememberme/pkg/router"
)

// Group is a Route factory that applies shared attributes to Routes created.
type Group struct {
	pathPrefix string
}

// SetPathPrefix sets the path prefix.
func (g *Group) SetPathPrefix(prefix string) {
	g.pathPrefix = fmt.Sprintf("/%v", strings.Trim(prefix, "/"))
}

// NewRoute creates a new Route instance using the group attributes.
func (g *Group) NewRoute(method, path string) *router.Route {
	if g.pathPrefix != "" {
		path = fmt.Sprintf("%v/%v", g.pathPrefix, strings.Trim(path, "/"))
	}

	return router.NewRoute(method, path)
}

// Get creates a new router.MethodGet Route using the group attributes.
func (g *Group) Get(path string) *router.Route {
	return g.NewRoute(router.MethodGet, path)
}

// Post creates a new router.MethodPost Route using the group attributes
func (g *Group) Post(path string) *router.Route {
	return g.NewRoute(router.MethodPost, path)
}

// Put creates a new router.MethodPut Route using the group attributes
func (g *Group) Put(path string) *router.Route {
	return g.NewRoute(router.MethodPut, path)
}

// Delete creates a new router.MethodDelete Route using the group attributes
func (g *Group) Delete(path string) *router.Route {
	return g.NewRoute(router.MethodDelete, path)
}

// NewGroup creates a new Group instance.
func NewGroup() *Group {
	return &Group{}
}
