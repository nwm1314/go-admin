package service

import (
	"go-admin/server/common"
	"go-admin/server/dao"
	"go-admin/server/model/response"
)

func GetApiListPage(page, pageSize int) common.Data {
	var apiPageResp response.ListPageResp
	apis, _ := dao.GetApiListByPage(page, pageSize)
	count := dao.GetAllApiCount()
	apiPageResp.Page = page
	apiPageResp.PageSize = pageSize
	apiPageResp.List = apis
	apiPageResp.Total = count
	return common.NewData(0, apiPageResp, "获取成功")
}
