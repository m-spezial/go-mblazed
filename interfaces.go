package mblazed

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)
type ICoreEngine interface {
	GetDb()
	GetRenderEngine() IRenderEngine
}
type IRequestContext interface {
	GetString(key string)string
	SetString(key string, value string)
	GetBool(key string) bool
	SetBool(key string, value bool)
	GetData(key string) interface{}
	SetData(key string, data interface{})
	GetResponseWriter() http.ResponseWriter
	GetRequest() *http.Request
	GetParams() httprouter.Params
}
