package service

import (
	"go-admin/server/common"
	"go-admin/server/dao"
	"go-admin/server/model/response"
)

func GetUserList(page, pageSize int) common.Data {
	var userPageResp response.ListPageResp
	ids, _ := dao.GetUserIdsByPage(page, pageSize)
	if ids != nil {
		var users []response.UserInfo
		for _, id := range ids {
			user := GetUserData(id)
			users = append(users, user)
		}
		userPageResp.List = users
	}
	userPageResp.Total = dao.GetAllUserCount()
	userPageResp.Page = page
	userPageResp.PageSize = pageSize
	return common.NewData(0, userPageResp, "获取成功")
}
