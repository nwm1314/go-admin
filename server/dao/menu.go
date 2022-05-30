package dao

import (
	"go-admin/server/model/database"
	"log"
)

func GetAllRootMenuId(page, pageSize int) ([]uint, error) {
	start := (page - 1) * pageSize
	var menuIds []uint
	err := DB.Table("sys_base_menus").Select("id").Where("parent_id=?", 0).Offset(start).Limit(pageSize).Scan(&menuIds).Error
	if err != nil {
		log.Printf("根据id获取子菜单id错误:%s", err.Error())
		return nil, err
	}
	return menuIds, nil
}

func GetMenuListByAuthorityId(authorityId string) ([]database.Menu, error) {
	var MenuIds []int
	var Menus []database.Menu
	err := DB.Where("id in (?)", DB.Table("sys_authority_menus").Where("sys_authority_authority_id=?", authorityId).Pluck("sys_base_menu_id", &MenuIds)).Find(&Menus).Error
	if err != nil {
		log.Printf("根据权限id获取菜单错误:%s", err.Error())
		return nil, err
	}
	return Menus, nil
}

func GetMenuIdListByAuthorityId(authorityId string) ([]uint, error) {
	var menuIds []uint
	err := DB.Table("sys_authority_menus").Select("sys_base_menu_id").Where("sys_authority_authority_id=?", authorityId).Scan(&menuIds).Error
	if err != nil {
		log.Printf("根据权限id获取菜单id错误:%s", err.Error())
		return nil, err
	}
	return menuIds, nil
}

func GetMenuById(menuId uint) (*database.Menu, error) {
	var menu database.Menu
	err := DB.Where("id=?", menuId).First(&menu).Error
	if err != nil {
		log.Printf("根据id获取菜单错误:%s", err.Error())
		return nil, err
	}
	return &menu, nil
}

func GetChildMenuByMenuId(menuId uint) ([]database.Menu, error) {
	var childMenus []database.Menu
	err := DB.Where("parent_id=?", menuId).Find(&childMenus).Error
	if err != nil {
		log.Printf("根据id获取子菜单错误:%s", err.Error())
		return nil, err
	}
	return childMenus, nil
}

func GetChildMenuIdByMenuId(menuId uint) ([]uint, error) {
	var childMenuIds []uint
	err := DB.Table("sys_base_menus").Select("id").Where("parent_id=?", menuId).Scan(&childMenuIds).Error
	if err != nil {
		log.Printf("根据id获取子菜单id错误:%s", err.Error())
		return nil, err
	}
	return childMenuIds, nil
}

func GetMenuBtnIdListByMenuId(menuId uint) ([]uint, error) {
	var menuBtnIds []uint
	err := DB.Table("sys_base_menu_btns").Select("id").Where("sys_base_menu_id=?", menuId).Scan(&menuBtnIds).Error
	if err != nil {
		log.Printf("根据菜单id获取菜单按钮id错误:%s", err.Error())
		return nil, err
	}
	return menuBtnIds, nil
}

func GetMenuBtnListByMenuId(menuId uint) ([]database.MenuBtn, error) {
	var menuBtns []database.MenuBtn
	err := DB.Where("sys_base_menu_id=?", menuId).Find(&menuBtns).Error
	if err != nil {
		log.Printf("获取菜单按钮错误:%s", err.Error())
		return nil, err
	}
	return menuBtns, nil
}

func GetMenuBtnById(id uint) (*database.MenuBtn, error) {
	var menuBtn database.MenuBtn
	err := DB.Where("id=?", id).First(&menuBtn).Error
	if err != nil {
		log.Printf("根据id获取菜单按钮错误:%s", err.Error())
		return nil, err
	}
	return &menuBtn, nil
}

func GetMenuParameterListByMenuId(menuId uint) ([]database.MenuParameter, error) {
	var menuParameters []database.MenuParameter
	err := DB.Where("sys_base_menu_id=?", menuId).Find(&menuParameters).Error
	if err != nil {
		log.Printf("获取菜单参数错误:%s", err.Error())
		return nil, err
	}
	return menuParameters, nil
}

func GetMenuParameterIdListByMenuId(menuId uint) ([]uint, error) {
	var menuParameterIds []uint
	err := DB.Table("sys_base_menu_parameters").Select("id").Where("sys_base_menu_id=?", menuId).Scan(&menuParameterIds).Error
	if err != nil {
		log.Printf("根据菜单id获取菜单参数id错误:%s", err.Error())
		return nil, err
	}
	return menuParameterIds, nil
}

func GetMenuParameterById(id uint) (*database.MenuParameter, error) {
	var menuParameter database.MenuParameter
	err := DB.Where("id=?", id).First(&menuParameter).Error
	if err != nil {
		log.Printf("根据id获取菜单参数错误:%s", err.Error())
		return nil, err
	}
	return &menuParameter, nil
}
