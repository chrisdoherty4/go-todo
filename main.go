package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/chrisdoherty4/rememberme/internal/pkg/todo"
)

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// var port = 8080

// var handlers = map[string]http.Handler{
// 	"/": rootHandler{},
// }

// func main() {
// 	for path, handler := range handlers {
// 		http.Handle(path, handler)
// 	}

// 	log.Printf("Starting server on %d", port)
// 	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
// }

func main() {
	item := todo.NewItem("Walk dog")

	marshalled, err := json.Marshal(item)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(marshalled))
}
