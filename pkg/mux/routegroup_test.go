package mux_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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
			func(_ http.ResponseWriter, _ *http.Request, _ *mux.RouteMatch) {},
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
			func(_ http.ResponseWriter, _ *http.Request, _ *mux.RouteMatch) {},
		))
	})

	assert.Equal(t, 1, router.Count())

	routes := router.RouteFromPath("/example")

	assert.Equal(t, 1, len(routes))
}

func TestGroupWithSimilarPaths(t *testing.T) {
	var (
		items              string = "items"
		itemsWithPathParam string = "itemsWithPathParam"
	)

	r.Group("/items", func(rg *mux.RouteGroup) {

		rg.Get("/", mux.NewInlineHandler(
			func(w http.ResponseWriter, r *http.Request, _ *mux.RouteMatch) {
				t.Log("Received request for /items")
				w.Write([]byte(items))
			},
		))

		rg.Get("/([A-Za-z-]+)", mux.NewInlineHandler(
			func(w http.ResponseWriter, r *http.Request, _ *mux.RouteMatch) {
				t.Log("Received request for /items/([A-Za-z-]+)")
				w.Write([]byte(itemsWithPathParam))
			},
		))

	})

	server := httptest.NewServer(r)
	defer server.Close()

	client := server.Client()

	testUrl, _ := newUrl(server.URL, "items", "walk-dog")
	response, _ := client.Get(testUrl.String())
	body, _ := ioutil.ReadAll(response.Body)

	t.Log("Body received was", string(body))
	assert.Equal(t, itemsWithPathParam, string(body))
}

func newUrl(parts ...string) (*url.URL, error) {
	return url.Parse(strings.Join(parts, "/"))
}
