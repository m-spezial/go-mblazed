package logic

type CoreEngine interface {
	GetDb()
	GetRenderEngine() RenderEngine
}

type RenderEngine interface {
	RenderHtml(templateName string)
}
