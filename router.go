package mblazed

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Router struct {
	httprouter.Router
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
