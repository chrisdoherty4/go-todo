package mux_test

import (
	"net/http"
	"testing"

	"github.com/chrisdoherty4/rememberme/pkg/mux"
	"github.com/stretchr/testify/assert"
)

func TestNewRouteGroup(t *testing.T) {
	router := mux.NewRouter()

	prefix := "/example"
	group := mux.NewRouteGroup(prefix, router)

	assert.Equal(t, prefix, group.PathPrefix())
}

func TestGroupWithSlashPrefix(t *testing.T) {
	router := mux.NewRouter()
	group := mux.NewRouteGroup("/", router)

	group.Group("/", func(rg *mux.RouteGroup) {
		rg.Get("/example", mux.NewInlineHandler(
			func(_ http.ResponseWriter, _ *http.Request) {},
		))
	})

	assert.Equal(t, 1, router.Count())

	routes := router.RouteFromPath("/example")

	assert.Equal(t, 1, len(routes))
}

func TestGroupWithWordPrefix(t *testing.T) {
	router := mux.NewRouter()
	group := mux.NewRouteGroup("/", router)

	group.Group("/example", func(rg *mux.RouteGroup) {
		rg.Get("/", mux.NewInlineHandler(
			func(_ http.ResponseWriter, _ *http.Request) {},
		))
	})

	assert.Equal(t, 1, router.Count())

	routes := router.RouteFromPath("/example")

	assert.Equal(t, 1, len(routes))
}
