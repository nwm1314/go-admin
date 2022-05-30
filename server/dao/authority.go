package dao

import (
	"go-admin/server/model/database"
	"log"
)

func GetAuthorityIdByUserId(userId uint) ([]string, error) {
	var authorityIds []string
	err := DB.Select("sys_authority_authority_id").Where("sys_user_id=?", userId).Scan(&authorityIds).Error
	if err != nil {
		log.Printf("获取用户权限id错误:%s", err.Error())
		return nil, err
	}
	return authorityIds, nil
}

func GetAuthorityById(id string) (*database.Authority, error) {
	var authority database.Authority
	err := DB.Where("authority_id=?", id).First(&authority).Error
	if err != nil {
		log.Printf("根据id查询权限错误:%s", err.Error())
		return nil, err
	}
	return &authority, nil
}

func GetAuthorityListByUserId(userId uint) ([]database.Authority, error) {
	var authorityIds []string
	var authorities []database.Authority
	err := DB.Where("id in (?)", DB.Table("sys_user_authority").Where("sys_user_id=?", userId).Pluck("sys_authority_authority_id", &authorityIds)).Find(&authorities).Error
	if err != nil {
		log.Printf("获取用户权限错误:%s", err.Error())
		return nil, err
	}
	return authorities, nil
}

func GetAuthorityIdListByUserId(userId uint) ([]string, error) {
	var authorityIds []string
	err := DB.Table("sys_user_authority").Select("sys_authority_authority_id").Where("sys_user_id=?", userId).Scan(&authorityIds).Error
	if err != nil {
		log.Printf("根据用户id获取权限id错误:%s", err.Error())
		return nil, err
	}
	return authorityIds, nil
}

func GetChildAuthorityListByAuthorityId(authorityId string) ([]database.Authority, error) {
	var authorities []database.Authority
	err := DB.Where("parent_id=?", authorityId).Find(&authorities).Error
	if err != nil {
		log.Printf("根据权限id获取子权限错误:%s", err.Error())
		return nil, err
	}
	return authorities, nil
}

func GetChildAuthorityIdListByAuthorityId(authorityId string) ([]string, error) {
	var authorityIds []string
	err := DB.Table("sys_authorities").Select("authority_id").Where("parent_id=?", authorityId).Scan(&authorityIds).Error
	if err != nil {
		log.Printf("根据权限id获取子权限id错误:%s", err.Error())
		return nil, err
	}
	return authorityIds, nil
}

func GetAuthorityIdListWithoutParentId() ([]string, error) {
	var authorityIds []string
	err := DB.Table("sys_authorities").Select("authority_id").Where("parent_id=?", 0).Scan(&authorityIds).Error
	if err != nil {
		log.Printf("查询无父权限id错误:%s", err.Error())
		return nil, err
	}
	return authorityIds, nil
}

func GetDataAuthorityListByAuthorityId(authorityId string) ([]database.Authority, error) {
	var authorities []database.Authority
	var dataAuthorityIds []string
	err := DB.Where("authority_id in (?)", DB.Table("sys_data_authority_id").Where("sys_authority_authority_id=?", authorityId).Pluck("data_authority_id_authority_id", &dataAuthorityIds)).Find(&authorities).Error
	if err != nil {
		log.Printf("根据权限id获取数据权限错误:%s", err.Error())
		return nil, err
	}
	return authorities, nil
}

func GetDataAuthorityIdListByAuthorityId(authorityId string) ([]string, error) {
	var dataAuthorityIds []string
	err := DB.Table("sys_data_authority_id").Select("data_authority_id_authority_id").Where("sys_authority_authority_id=?", authorityId).Scan(&dataAuthorityIds).Error
	if err != nil {
		log.Printf("根据权限id获取数据权限id错误:%s", err.Error())
		return nil, err
	}
	return dataAuthorityIds, nil
}

func GetAuthorityListByMenuId(menuId uint) ([]database.Authority, error) {
	var authorities []database.Authority
	var authorityIds []string
	err := DB.Where("authority_id in (?)", DB.Table("sys_authority_menus").Where("sys_base_menu_id=?", menuId).Pluck("sys_authority_authority_id", &authorityIds)).Find(&authorities).Error
	if err != nil {
		log.Printf("根据菜单id获取权限列表错误:%s", err.Error())
		return nil, err
	}
	return authorities, nil
}

func GetAuthorityIdListByMenuId(menuId uint) ([]string, error) {
	var authorityIds []string
	err := DB.Table("sys_authority_menus").Select("sys_authority_authority_id").Where("sys_base_menu_id=?", menuId).Scan(&authorityIds).Error
	if err != nil {
		log.Printf("根据菜单id获取权限id错误:%s", err.Error())
		return nil, err
	}
	return authorityIds, nil
}
