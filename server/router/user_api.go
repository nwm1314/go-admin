package router

import "go-admin/server/api"

type UserApi struct {
}

func (UserApi) InitUserRouter() {
	userApiRouter := Router.Group("/user")
	userApi := &api.UserApi{}
	{
		userApiRouter.GET("/getUserInfo", userApi.GetUserInfo)
		userApiRouter.POST("/getUserList", userApi.GetUserList)
	}
}
