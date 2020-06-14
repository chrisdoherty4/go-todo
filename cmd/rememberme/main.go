package main

import (
	"fmt"
	"log"
	"net/http"
)

var port = 8080

func main() {
	log.Printf("Starting server on %d", port)

	router := NewRouter()

	router.HandleInline(GetRoute("/"), func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Inline handler"))
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
