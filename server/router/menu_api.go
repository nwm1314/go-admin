package router

import (
	"go-admin/server/api"
)

type MenuApi struct {
}

func (*MenuApi) InitMenuRouter() {
	menuApiRouter := Router.Group("/menu")
	menuApi := &api.MenuApi{}
	{
		menuApiRouter.POST("/getMenu", menuApi.HandleGetMenu)
		menuApiRouter.POST("/getMenuList", menuApi.HandleGetMenuList)
	}
}
