package logic

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type RequestContextInterface interface {
	GetData(key string) interface{}
	SetData(key string, data interface{})
	GetResponseWriter() http.ResponseWriter
	GetRequest() *http.Request
	GetParams() httprouter.Params
}
