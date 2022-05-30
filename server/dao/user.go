package dao

import (
	"go-admin/server/model/database"
	"log"
)

func GetUserByUsername(username string) (*database.User, error) {
	var user database.User
	err := DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		log.Printf("根据用户名获取用户错误:%s", err.Error())
		return nil, err
	}
	return &user, nil
}

func GetUserById(id uint) (*database.User, error) {
	var user database.User
	err := DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Printf("根据id获取用户错误:%s", err.Error())
		return nil, err
	}
	return &user, nil
}

func GetUserIdsByPage(page, pageSize int) ([]uint, error) {
	var ids []uint
	start := (page - 1) * pageSize
	err := DB.Table("sys_users").Select("id").Offset(start).Limit(pageSize).Scan(&ids).Error
	if err != nil {
		log.Printf("分页获取用户id错误:%s", err.Error())
		return nil, err
	}
	return ids, nil
}

func GetAllUserCount() int64 {
	var count int64
	err := DB.Table("sys_users").Count(&count).Error
	if err != nil {
		log.Printf("获取用户数量错误:%s", err.Error())
		return 0
	}
	return count
}
