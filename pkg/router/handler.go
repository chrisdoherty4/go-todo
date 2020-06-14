package router

import "net/http"

// Handler defines an interface to handle requests.
type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

// InlineHandler is a type for writing an inline function handler.
type InlineHandler func(http.ResponseWriter, *http.Request)

// inlineHandlerWrapper wraps up an InlineHandler type (a function signature)
// so that it can be used with a RouteHandler.
type inlineHandlerWrapper struct {
	handler InlineHandler
}

func (t inlineHandlerWrapper) Handle(w http.ResponseWriter, r *http.Request) {
	t.handler(w, r)
}
