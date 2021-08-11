package mblazed

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Router struct {
	httprouter.Router
}

func wrapRequestHandle(handle RequestHandler) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		ctx := NewRequestContext(writer, request, params)
		handle(ctx)
	}
}

func NewRouter() *Router  {
	return &Router{httprouter.Router{
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      true,
		HandleMethodNotAllowed: true,
		HandleOPTIONS:          true,
		GlobalOPTIONS:          nil,
		NotFound:               nil,
		MethodNotAllowed:       nil,
		PanicHandler:           nil,
	}}
}
