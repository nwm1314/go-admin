package service

import (
	"go-admin/server/common"
	"go-admin/server/dao"
	"go-admin/server/model/response"
	"strconv"
)

func GetMenuList(page, pageSize int) common.Data {
	var menuListResp response.ListPageResp
	menuListResp.Page = page
	menuListResp.PageSize = pageSize
	rootIds, _ := dao.GetAllRootMenuId(page, pageSize)
	menuListResp.Total = int64(len(rootIds))
	if rootIds != nil {
		var rootMenus []response.Menu
		for _, rootId := range rootIds {
			rootMenu := GetMenuDataById(rootId)
			rootMenus = append(rootMenus, rootMenu)
		}
		menuListResp.List = rootMenus
	}
	return common.NewData(0, menuListResp, "获取成功")
}

func GetMenu(authorityId string) common.Data {
	menuIds, _ := dao.GetMenuIdListByAuthorityId(authorityId)
	if menuIds != nil {
		var menuResp response.MenuResp
		for _, menuId := range menuIds {
			menu := GetMenuDataById(menuId)
			menuResp.Menus = append(menuResp.Menus, menu)
		}
		return common.NewData(0, menuResp, "获取成功")
	}
	return common.NewData(0, nil, "暂无数据")
}

func GetMenuBtnDataById(id uint) response.MenuBtb {
	var menuBtnData response.MenuBtb
	menuBtn, _ := dao.GetMenuBtnById(id)
	menuBtnData.SysBaseMenuID = menuBtn.SysBaseMenuId
	menuBtnData.Id = menuBtn.ID
	menuBtnData.Name = menuBtn.Name
	menuBtnData.UpdatedAt = menuBtn.UpdatedAt.String()
	menuBtnData.CreatedAt = menuBtn.CreatedAt.String()
	menuBtnData.Desc = menuBtn.Desc
	return menuBtnData
}

func GetMenuParameterDataById(id uint) response.MenuParameter {
	var menuParameterData response.MenuParameter
	menuParameter, _ := dao.GetMenuParameterById(id)
	menuParameterData.SysBaseMenuID = menuParameter.SysBaseMenuId
	menuParameterData.Id = menuParameter.ID
	menuParameterData.CreatedAt = menuParameter.CreatedAt.String()
	menuParameterData.UpdatedAt = menuParameter.UpdatedAt.String()
	menuParameterData.Type = menuParameter.Type
	menuParameterData.Key = menuParameter.Key
	menuParameterData.Value = menuParameter.Value
	return menuParameterData
}

func GetMenuDataById(menuId uint) response.Menu {
	var menuData response.Menu
	menu, _ := dao.GetMenuById(menuId)
	menuData.Id = menu.ID
	menuData.MenuId = strconv.Itoa(int(menu.ID))
	menuData.UpdatedAt = menu.UpdatedAt.String()
	menuData.ParentId = menu.ParentId
	menuData.CreatedAt = menu.CreatedAt.String()
	menuData.Path = menu.Path
	menuData.Name = menu.Name
	menuData.Component = menu.Component
	menuData.Sort = menu.Sort
	menuData.Hidden = false
	if menu.Hidden == 1 {
		menuData.Hidden = true
	}
	var meta response.Meta
	meta.DefaultMenu = false
	if menu.DefaultMenu == 1 {
		meta.DefaultMenu = true
	}
	meta.Title = menu.Title
	meta.KeepAlive = false
	if menu.KeepAlive == 1 {
		meta.KeepAlive = true
	}
	meta.CloseTab = false
	if menu.CloseTab == 1 {
		meta.CloseTab = true
	}
	meta.Icon = menu.Icon
	menuData.Meta = meta

	childMenuIds, _ := dao.GetChildMenuIdByMenuId(menuId)
	if childMenuIds != nil {
		var childMenus []response.Menu
		for _, childMenuId := range childMenuIds {
			childMenu := GetMenuDataById(childMenuId)
			childMenus = append(childMenus, childMenu)
		}
		menuData.Children = childMenus
	}
	menuBtnIds, _ := dao.GetMenuBtnIdListByMenuId(menuId)
	if menuBtnIds != nil {
		var menuBtns []response.MenuBtb
		for _, menuBtnId := range menuBtnIds {
			menuBtn := GetMenuBtnDataById(menuBtnId)
			menuBtns = append(menuBtns, menuBtn)
		}
		menuData.MenuBtn = menuBtns
	}
	menuParameterIds, _ := dao.GetMenuParameterIdListByMenuId(menuId)
	if menuParameterIds != nil {
		var menuParameters []response.MenuParameter
		for _, menuParameterId := range menuParameterIds {
			menuParameter := GetMenuParameterDataById(menuParameterId)
			menuParameters = append(menuParameters, menuParameter)
		}
		menuData.Parameters = menuParameters
	}
	return menuData
}
