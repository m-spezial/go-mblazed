package mblazed

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func DefaultRouter() *gin.Engine {
	router := gin.Default()
	router.Use(sessions.Sessions("GoldenYarn",cookie.NewStore([]byte("secret"))))
	return router
}
