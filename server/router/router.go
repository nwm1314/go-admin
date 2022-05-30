package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/server/middleware"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	Router.Use(middleware.TokenMiddle())
	Router.SetTrustedProxies(nil)
}
