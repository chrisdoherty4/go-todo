package route_test

import (
	"testing"

	"github.com/chrisdoherty4/rememberme/pkg/router/route"
	"github.com/stretchr/testify/assert"
)

var group = route.NewGroup()

func init() {
	group.SetPathPrefix("/api")
}

func TestCreatingGetRoute(t *testing.T) {
	r := group.Get("/get")
	assert.Equal(t, "/api/get", r.Path)
}

func TestCreatingPostRoute(t *testing.T) {
	r := group.Post("/post")
	assert.Equal(t, "/api/post", r.Path)
}

func TestCreatingPutRoute(t *testing.T) {
	r := group.Put("/put")
	assert.Equal(t, "/api/put", r.Path)
}

func TestCreatingDeleteRoute(t *testing.T) {
	r := group.Delete("/delete")
	assert.Equal(t, "/api/delete", r.Path)
}

func TestingExcessiveSlashesInPathPrefix(t *testing.T) {
	group := route.NewGroup()
	group.SetPathPrefix("////api////")

	r := group.Get("/post")

	assert.Equal(t, "/api/post", r.Path)
}
