package mblazed

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ICoreEngine interface {
	GetDb()
	GetRenderEngine() IRenderEngine
}

type IRenderEngine interface {
	RenderHtml(templateName string)
}

type IRequestContext interface {
	GetData(key string) interface{}
	SetData(key string, data interface{})
	GetResponseWriter() http.ResponseWriter
	GetRequest() *http.Request
	GetParams() httprouter.Params
}

type IMblazedPlugin interface {
	
}
