package service

import (
	"go-admin/server/common"
	"go-admin/server/dao"
	"go-admin/server/model/response"
)

func GetAuthorityList(page, pageSize int) common.Data {
	var authorityListResp response.ListPageResp
	authorityListResp.Page = page
	authorityListResp.PageSize = pageSize
	authorityIds, _ := dao.GetAuthorityIdListWithoutParentId()
	authorityListResp.Total = int64(len(authorityIds))
	if authorityIds != nil {
		var authorities []response.Authority
		for _, authorityId := range authorityIds {
			authority := GetAuthorityDataById(authorityId)
			authorities = append(authorities, authority)
		}
		authorityListResp.List = authorities
	}
	return common.NewData(0, authorityListResp, "获取成功")
}
