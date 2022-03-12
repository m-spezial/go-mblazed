package core

import (
	"code.m-spezial.de/M-Spezial/go-mblazed/logic"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Engine struct {
	router                       *Router
	requestContextProcessorChain []logic.RequestContextProcessor
}

func NewEngine() logic.CoreEngine {
	return &Engine{
		router: NewRouter(),
	}
}

func (ce *Engine) GetDb() {
	//TODO implement me
	panic("implement me")
}

func (ce *Engine) GetRenderEngine() logic.RenderEngine {
	//TODO implement me
	panic("implement me")
}

func (ce *Engine) wrapRequestHandle(handle logic.RequestHandler) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var ctx logic.RequestContextInterface
		ctx = NewRequestContext(writer, request, params)
		for i := 0; i < len(ce.requestContextProcessorChain); i++ {
			ctx = ce.requestContextProcessorChain[i](ctx)
		}
		handle(ctx)
	}
}

// GET is a shortcut for ce.Handle(http.MethodGet, path, handle)
func (ce *Engine) GET(pathname string, path string, handle logic.RequestHandler) {
	ce.Handle(http.MethodGet, pathname, path, handle)
}

// HEAD is a shortcut for ce.Handle(http.MethodHead, path, handle)
func (ce *Engine) HEAD(pathname string, path string, handle logic.RequestHandler) {
	ce.Handle(http.MethodHead, pathname, path, handle)
}

// OPTIONS is a shortcut for ce.Handle(http.MethodOptions, path, handle)
func (ce *Engine) OPTIONS(pathname string, path string, handle logic.RequestHandler) {
	ce.Handle(http.MethodOptions, pathname, path, handle)
}

// POST is a shortcut for ce.Handle(http.MethodPost, path, handle)
func (ce *Engine) POST(pathname string, path string, handle logic.RequestHandler) {
	ce.Handle(http.MethodPost, pathname, path, handle)
}

// PUT is a shortcut for ce.Handle(http.MethodPut, path, handle)
func (ce *Engine) PUT(pathname string, path string, handle logic.RequestHandler) {
	ce.Handle(http.MethodPut, pathname, path, handle)
}

// PATCH is a shortcut for ce.Handle(http.MethodPatch, path, handle)
func (ce *Engine) PATCH(pathname string, path string, handle logic.RequestHandler) {
	ce.Handle(http.MethodPatch, pathname, path, handle)
}

// DELETE is a shortcut for ce.Handle(http.MethodDelete, path, handle)
func (ce *Engine) DELETE(pathname string, path string, handle logic.RequestHandler) {
	ce.Handle(http.MethodDelete, pathname, path, handle)
}

func (ce *Engine) Handle(methode string, pathname string, path string, handle logic.RequestHandler) {
	ce.router.Handle(methode, pathname, path, ce.wrapRequestHandle(handle))
}

func (ce *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("[mblazed] incoming request:" + request.URL.Path)
	ce.router.ServeHTTP(writer, request)
}
