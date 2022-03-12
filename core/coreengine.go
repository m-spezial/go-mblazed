package core

import (
	"code.m-spezial.de/M-Spezial/go-mblazed/logic"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type RequestHandler func(r logic.RequestContextInterface)
type RequestContextProcessor func(ctx logic.RequestContextInterface) logic.RequestContextInterface

type coreEngine struct {
	router                       *Router
	requestContextProcessorChain []RequestContextProcessor
}

func NewCoreEngine() logic.CoreEngine {
	return &coreEngine{
		router: NewRouter(),
	}
}

func (ce coreEngine) GetDb() {
	//TODO implement me
	panic("implement me")
}

func (ce coreEngine) GetRenderEngine() logic.RenderEngine {
	//TODO implement me
	panic("implement me")
}

func (ce *coreEngine) wrapRequestHandle(handle RequestHandler) httprouter.Handle {
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
func (ce *coreEngine) GET(pathname string, path string, handle RequestHandler) {
	ce.Handle(http.MethodGet, pathname, path, handle)
}

// HEAD is a shortcut for ce.Handle(http.MethodHead, path, handle)
func (ce *coreEngine) HEAD(pathname string, path string, handle RequestHandler) {
	ce.Handle(http.MethodHead, pathname, path, handle)
}

// OPTIONS is a shortcut for ce.Handle(http.MethodOptions, path, handle)
func (ce *coreEngine) OPTIONS(pathname string, path string, handle RequestHandler) {
	ce.Handle(http.MethodOptions, pathname, path, handle)
}

// POST is a shortcut for ce.Handle(http.MethodPost, path, handle)
func (ce *coreEngine) POST(pathname string, path string, handle RequestHandler) {
	ce.Handle(http.MethodPost, pathname, path, handle)
}

// PUT is a shortcut for ce.Handle(http.MethodPut, path, handle)
func (ce *coreEngine) PUT(pathname string, path string, handle RequestHandler) {
	ce.Handle(http.MethodPut, pathname, path, handle)
}

// PATCH is a shortcut for ce.Handle(http.MethodPatch, path, handle)
func (ce *coreEngine) PATCH(pathname string, path string, handle RequestHandler) {
	ce.Handle(http.MethodPatch, pathname, path, handle)
}

// DELETE is a shortcut for ce.Handle(http.MethodDelete, path, handle)
func (ce *coreEngine) DELETE(pathname string, path string, handle RequestHandler) {
	ce.Handle(http.MethodDelete, pathname, path, handle)
}

func (ce *coreEngine) Handle(methode string, pathname string, path string, handle RequestHandler) {
	ce.router.Handle(methode, pathname, path, ce.wrapRequestHandle(handle))
}

func (ce *coreEngine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("[mblazed] incoming request:" + request.URL.Path)
	ce.router.ServeHTTP(writer, request)
}
