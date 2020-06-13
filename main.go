package main

import (
	"fmt"
	"log"
	"net/http"
)

var port = 8080

var handlers = map[string]http.Handler{
	"/": rootHandler{},
}

func main() {
	for path, handler := range handlers {
		http.Handle(path, handler)
	}

	log.Printf("Starting server on %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
