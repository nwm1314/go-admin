package main

import "go-admin/server/router"

func main() {
	var baseApi = &router.BaseApi{}
	var menuApi = &router.MenuApi{}
	var userApi = &router.UserApi{}
	var authorityApi = &router.AuthorityApi{}
	var apiApi = &router.ApiApi{}
	baseApi.InitBaseRouter()
	menuApi.InitMenuRouter()
	userApi.InitUserRouter()
	authorityApi.InitAuthorityRouter()
	apiApi.InitApiRouter()
	router.Router.Run(":8085")
}
