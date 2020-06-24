package mux_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chrisdoherty4/rememberme/pkg/mux"
	"github.com/stretchr/testify/assert"
)

func TestServeingRequests(t *testing.T) {
	expected := "Replying to /example request"

	r := mux.NewRouter()

	r.Get("/example", mux.NewInlineHandler(
		func(w http.ResponseWriter, r *http.Request, _ *mux.RouteMatch) {
			t.Log("Received request for /example")
			w.Write([]byte(expected))
		},
	))

	request, err := http.NewRequest(http.MethodGet, "/example", nil)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, request)

	body, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	assert.Equal(t, expected, string(body))
}
