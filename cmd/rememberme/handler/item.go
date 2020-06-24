package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/chrisdoherty4/rememberme/pkg/mux"
	"github.com/chrisdoherty4/rememberme/pkg/todo"
)

// ItemController handles CRUD interactions with a todo repository.
type ItemController struct {
	repo todo.Repository
}

// List lists all items in the todo repository.
func (ic *ItemController) List(w http.ResponseWriter, r *http.Request, rm *mux.RouteMatch) {
	// This is dodgey because there could be a huge number of items in the
	// repository. Think about capping this somehow and introducing an offset
	// and size to the Repository.GetAll() method so we can control how many
	// things are returned.
	//
	// Could do with just an offset and leave the caller, here, to decide how
	// many are read out.
	items := make([]todo.Item, ic.repo.Size())

	i := 0
	for item := range ic.repo.GetAll() {
		items[i] = item
		i++
	}

	marshalledItems, err := json.MarshalIndent(items, "", "  ")

	if err != nil {
		log.Printf("Attempted marshalling todo items: \n %v", err.Error())
		ServerError(w, r)
		return
	}

	io.WriteString(w, string(marshalledItems))
}

// Show retrieves a single item from the repository and returns it as json.
func (ic *ItemController) Show(w http.ResponseWriter, r *http.Request, rm *mux.RouteMatch) {
	title, _ := rm.Var(0)

	item, err := ic.repo.Get(title)
	if err != nil {
		log.Printf("Requested item not found: %v", r.URL.String())
		NotFound(w, r)
		return
	}

	marshalledItem, err := json.MarshalIndent(item, "", "  ")

	if err != nil {
		log.Printf("Failed json marshalling: \n %v", err.Error())
		ServerError(w, r)
		return
	}

	io.WriteString(w, string(marshalledItem))
}

// Save saves a new todo item.
func (ic *ItemController) Save(w http.ResponseWriter, r *http.Request, rm *mux.RouteMatch) {
	var item todo.Item

	title, err := rm.Var(0)
	if err != nil {
		log.Printf("Couldn't read title from URL: \n %v", err.Error())
		ServerError(w, r)
		return
	}

	item.SetTitle(title)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read body of request: \n %v", err.Error())
		ServerError(w, r)
		return
	}

	if err := json.Unmarshal(body, &item); err != nil {
		log.Printf("Failed json unmarshalling: \n %v", err.Error())
		ServerError(w, r)
		return
	}

	ic.repo.Save(&item)
	io.WriteString(w, "Success")
}

// NewItemController creates a new ItemController instance.
func NewItemController(repo todo.Repository) *ItemController {
	return &ItemController{
		repo: repo,
	}
}
