package router

import "go-admin/server/api"

type ApiApi struct {
}

func (ApiApi) InitApiRouter() {
	apiApiRouter := Router.Group("/api")
	apiApi := &api.ApiApi{}
	{
		apiApiRouter.POST("/getApiList", apiApi.HandleApiList)
	}
}
