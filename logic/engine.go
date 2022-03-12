package logic

import "net/http"

type CoreEngine interface {
	HEAD(pathname string, path string, handle RequestHandler)
	GET(pathname string, path string, handle RequestHandler)
	PUT(pathname string, path string, handle RequestHandler)
	POST(pathname string, path string, handle RequestHandler)
	PATCH(pathname string, path string, handle RequestHandler)
	DELETE(pathname string, path string, handle RequestHandler)
	OPTIONS(pathname string, path string, handle RequestHandler)
	http.Handler
	GetDb()
	GetRenderEngine() RenderEngine
}

type RenderEngine interface {
	RenderHtml(templateName string)
}
