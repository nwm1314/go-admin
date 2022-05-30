package router

import "go-admin/server/api"

type AuthorityApi struct {
}

func (*AuthorityApi) InitAuthorityRouter() {
	authorityApiRouter := Router.Group("/authority")
	authorityApi := &api.AuthorityApi{}
	{
		authorityApiRouter.POST("/getAuthorityList", authorityApi.GetAuthorityList)
	}
}
