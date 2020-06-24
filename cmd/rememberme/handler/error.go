package handler

import (
	"fmt"
	"net/http"
)

// ServerError writes a generic server error and sets status 500
func ServerError(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "Something went wrong our end", http.StatusInternalServerError)
}

// NotFound writes a generic not found error and sets status 404.
func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("Not found %v", r.URL.String()), http.StatusNotFound)
}
