package main

import "net/http"

type rootHandler struct{}

func (t rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Root handler"))
}
