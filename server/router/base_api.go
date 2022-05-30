package router

import (
	"go-admin/server/api"
)

type BaseApi struct {
}

func (*BaseApi) InitBaseRouter() {
	baseApiRouter := Router.Group("/base")
	baseApi := &api.BaseApi{}
	{
		baseApiRouter.POST("/captcha", baseApi.HandleCaptcha)
		baseApiRouter.POST("/login", baseApi.HandleLogin)
	}
}
