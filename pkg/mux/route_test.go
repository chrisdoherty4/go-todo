package mux_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/chrisdoherty4/rememberme/pkg/mux"
	"github.com/stretchr/testify/assert"
)

var handler = mux.NewInlineHandler(func(_ http.ResponseWriter, _ *http.Request, _ *mux.RouteMatch) {})

func TestMatchBasicRoute(t *testing.T) {
	route := mux.NewRoute(mux.MethodGet, "/example", handler)

	assert.True(t, route.Match(newRequest("/example")))
}

func TestMatchRegexRoute(t *testing.T) {
	route := mux.NewRoute(mux.MethodGet, "/example/([0-9]+)", handler)

	assert.True(t, route.Match(newRequest("/example/1")))
}

func TestHandlingRegexRoute(t *testing.T) {
	route := mux.NewRoute(
		mux.MethodGet,
		"/example/([0-9]+)",
		mux.NewInlineHandler(
			func(w http.ResponseWriter, _ *http.Request, rm *mux.RouteMatch) {
				id, _ := rm.Var(0)
				w.Write([]byte(id))
			},
		),
	)

	recorder := httptest.NewRecorder()
	route.Handle(recorder, newRequest("/example/12345"))

	assert.Equal(t, "12345", recorder.Body.String())
}

func newRequest(path string) *http.Request {
	return httptest.NewRequest(
		http.MethodGet,
		fmt.Sprintf("http://localhost%v", path),
		strings.NewReader(""),
	)
}
