package mux_test

import (
	"net/http"
	"testing"

	"github.com/chrisdoherty4/rememberme/pkg/mux"
	"github.com/stretchr/testify/assert"
)

func TestGroupWithPathPrefix(t *testing.T) {
	prefix := "/example"
	group := mux.NewRouteGroup(prefix, mux.NewRouter())
	assert.Equal(t, prefix, group.PathPrefix())
}

func TestGroupWithSlashPrefix(t *testing.T) {
	r := mux.NewRouter()
	g := mux.NewRouteGroup("/", r)

	var expected *mux.Route

	g.Group("/", func(rf mux.RouteFactory) {
		expected = rf.Get("/example", mux.NewInlineHandler(
			func(_ http.ResponseWriter, _ *http.Request, _ *mux.RouteMatch) {},
		))
	})

	assert.Equal(t, 1, r.Count())

	request, err := http.NewRequest(http.MethodGet, "/example", nil)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	found := r.FindRoute(request)

	assert.NotNil(t, found)
	assert.True(t, found.Equal(expected))
}

func TestGroupWithWordPrefix(t *testing.T) {
	r := mux.NewRouter()
	g := mux.NewRouteGroup("/", r)

	var expected *mux.Route

	g.Group("/example", func(rf mux.RouteFactory) {
		expected = rf.Get("/", mux.NewInlineHandler(
			func(_ http.ResponseWriter, _ *http.Request, _ *mux.RouteMatch) {},
		))
	})

	assert.Equal(t, 1, r.Count())

	request, err := http.NewRequest(http.MethodGet, "/example", nil)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	found := r.FindRoute(request)

	assert.NotNil(t, found)
	assert.True(t, found.Equal(expected))
}

func TestGroupWithSimilarPaths(t *testing.T) {
	var expected *mux.Route

	r := mux.NewRouter()
	g := mux.NewRouteGroup("/", r)

	g.Group("/example", func(rf mux.RouteFactory) {

		rf.Get("/", mux.NewInlineHandler(
			func(_ http.ResponseWriter, _ *http.Request, _ *mux.RouteMatch) {
				t.Log("Received request for /example")
			},
		))

		expected = rf.Get("/([A-Za-z-]+)", mux.NewInlineHandler(
			func(_ http.ResponseWriter, _ *http.Request, _ *mux.RouteMatch) {
				t.Log("Received request for /example/([A-Za-z-])")
			},
		))

	})

	request, err := http.NewRequest(http.MethodGet, "/example/param", nil)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	found := r.FindRoute(request)

	assert.NotNil(t, found)
	assert.True(t, found.Equal(expected))
}

func TestPathTokenExtraction(t *testing.T) {

}

func TestRequestPathWithTrailingSlash(t *testing.T) {
	r := mux.NewRouter()
	g := mux.NewRouteGroup("/", r)

	expected := g.Get("/example", mux.NewInlineHandler(
		func(w http.ResponseWriter, r *http.Request, _ *mux.RouteMatch) {
			t.Log("Received request for /example")
		},
	))

	request, err := http.NewRequest(http.MethodGet, "/example/", nil)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	found := r.FindRoute(request)

	assert.NotNil(t, found)
	assert.True(t, expected.Equal(found))
}
