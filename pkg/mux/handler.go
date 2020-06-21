package mux

import "net/http"

// Handler defines an interface to handle requests.
type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

// HandlerFunc is a type for writing an inline function handler.
type HandlerFunc func(http.ResponseWriter, *http.Request)

// InlineHandler wraps up an InlineHandler type (a function signature)
// so that it can be used with a RouteHandler.
type InlineHandler struct {
	handler HandlerFunc
}

// Handle executes the HandlerFunc.
func (t InlineHandler) Handle(w http.ResponseWriter, r *http.Request) {
	t.handler(w, r)
}

// NewInlineHandler creates an instance of an InlineHandler that will execute
// the HandlerFunc during the `Handle()` call.
func NewInlineHandler(f HandlerFunc) *InlineHandler {
	return &InlineHandler{
		handler: f,
	}
}
