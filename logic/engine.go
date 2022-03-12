package logic

type CoreEngineInterface interface {
	GetDb()
	GetRenderEngine() RenderEngineInterface
}

type RenderEngineInterface interface {
	RenderHtml(templateName string)
}
