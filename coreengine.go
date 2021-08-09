package mblazed

import (
	"github.com/julienschmidt/httprouter"
	"log"
)
import "net/http"

type RequestHandler func(r IRequestContext)

type CoreEngine struct {
	router *httprouter.Router
	requestHandlerChain map[string] RequestHandler
}

func (ce CoreEngine) GetDB() {

}

func NewCoreEngine() *CoreEngine  {
	return  &CoreEngine{
		router: httprouter.New(),
	}
}

func wrapViewHandle(handle RequestHandler) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		ctx := NewRequestContext(writer, request, params)
		handle(ctx)
	}
}

// GET is a shortcut for ce.Handle(http.MethodGet, path, handle)
func (ce *CoreEngine) GET(path string, handle RequestHandler) {
	ce.Handle(http.MethodGet, path, handle)
}

// HEAD is a shortcut for ce.Handle(http.MethodHead, path, handle)
func (ce *CoreEngine) HEAD(path string, handle RequestHandler) {
	ce.Handle(http.MethodHead, path, handle)
}

// OPTIONS is a shortcut for ce.Handle(http.MethodOptions, path, handle)
func (ce *CoreEngine) OPTIONS(path string, handle RequestHandler) {
	ce.Handle(http.MethodOptions, path, handle)
}

// POST is a shortcut for ce.Handle(http.MethodPost, path, handle)
func (ce *CoreEngine) POST(path string, handle RequestHandler) {
	ce.Handle(http.MethodPost, path, handle)
}

// PUT is a shortcut for ce.Handle(http.MethodPut, path, handle)
func (ce *CoreEngine) PUT(path string, handle RequestHandler) {
	ce.Handle(http.MethodPut, path, handle)
}

// PATCH is a shortcut for ce.Handle(http.MethodPatch, path, handle)
func (ce *CoreEngine) PATCH(path string, handle RequestHandler) {
	ce.Handle(http.MethodPatch, path, handle)
}

// DELETE is a shortcut for ce.Handle(http.MethodDelete, path, handle)
func (ce *CoreEngine) DELETE(path string, handle RequestHandler) {
	ce.Handle(http.MethodDelete, path, handle)
}

func (ce *CoreEngine) Handle(methode string, path string, handle RequestHandler) {
	ce.router.Handle(methode, path, wrapViewHandle(handle))
}

func (ce *CoreEngine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("[mblazed] incoming request:" + request.URL.Path)
	ce.router.ServeHTTP(writer, request)
}
