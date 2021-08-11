package mblazed

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type RequestContext struct {
	PostData map[string] string
	stringData map[string] string
	boolData map[string] bool
	data map[string] interface{}
	writer http.ResponseWriter
	request *http.Request
	params httprouter.Params
}

func NewRequestContext(writer http.ResponseWriter, request *http.Request, params httprouter.Params) *RequestContext {
	return &RequestContext{
		PostData:   make(map[string] string),
		stringData: make(map[string] string),
		boolData:   make(map[string] bool),
		data:       make(map[string] interface{}),
		writer: 	writer,
		request: 	request,
		params: 	params,
	}
}

func (r *RequestContext) GetData(key string) interface{} {
	return r.data[key]
}

func (r *RequestContext) SetData(key string, data interface{}) {
	r.data[key] = data
}

func (r *RequestContext) GetResponseWriter() http.ResponseWriter {
	return  r.writer
}

func (r *RequestContext) GetRequest() *http.Request {
	return r.request
}

func (r *RequestContext) GetParams() httprouter.Params {
	return r.params
}
