package mblazed

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Router struct {
	httprouter.Router
	reverseMap map[string] string
}

func NewRouter() *Router  {
	return &Router{
		Router: httprouter.Router{
			RedirectTrailingSlash:  true,
			RedirectFixedPath:      true,
			HandleMethodNotAllowed: true,
			HandleOPTIONS:          true,
			GlobalOPTIONS:          nil,
			NotFound:               nil,
			MethodNotAllowed:       nil,
			PanicHandler:           nil,
		},
		reverseMap: make(map[string] string),
	}
}

// GET is a shortcut for r.Handle(http.MethodGet, path, handle)
func (r *Router) GET(pathname string, path string, handle httprouter.Handle) {
	r.Handle(http.MethodGet, pathname, path, handle)
}

// HEAD is a shortcut for r.Handle(http.MethodHead, path, handle)
func (r *Router) HEAD(pathname string, path string, handle httprouter.Handle) {
	r.Handle(http.MethodHead, pathname, path, handle)
}

// OPTIONS is a shortcut for r.Handle(http.MethodOptions, path, handle)
func (r *Router) OPTIONS(pathname string, path string, handle httprouter.Handle) {
	r.Handle(http.MethodOptions, pathname, path, handle)
}

// POST is a shortcut for r.Handle(http.MethodPost, path, handle)
func (r *Router) POST(pathname string, path string, handle httprouter.Handle) {
	r.Handle(http.MethodPost, pathname, path, handle)
}

// PUT is a shortcut for r.Handle(http.MethodPut, path, handle)
func (r *Router) PUT(pathname string, path string, handle httprouter.Handle) {
	r.Handle(http.MethodPut, pathname, path, handle)
}

// PATCH is a shortcut for r.Handle(http.MethodPatch, path, handle)
func (r *Router) PATCH(pathname string, path string, handle httprouter.Handle) {
	r.Handle(http.MethodPatch, pathname, path, handle)
}

// DELETE is a shortcut for r.Handle(http.MethodDelete, path, handle)
func (r *Router) DELETE(pathname string, path string, handle httprouter.Handle) {
	r.Handle(http.MethodDelete, pathname, path, handle)
}

func (r *Router) Handle(methode string, pathname string, path string, handle httprouter.Handle) {
	r.RegisterReverseRoute(pathname, path)
	r.Router.Handle(methode, path, handle)
}

func (r *Router) ANY (pathname string, path string, handle httprouter.Handle)  {
	r.RegisterReverseRoute(pathname, path)
}

func (r *Router) RegisterReverseRoute(name string, path string)  {
	r.reverseMap[name] = path
}
