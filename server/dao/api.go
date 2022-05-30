package dao

import (
	"go-admin/server/model/database"
	"log"
)

func GetApiListByPage(page, pageSize int) ([]database.Api, error) {
	var apis []database.Api
	start := (page - 1) * pageSize
	err := DB.Offset(start).Limit(pageSize).Find(&apis).Error
	if err != nil {
		log.Printf("分页查询api错误:%s", err.Error())
		return nil, err
	}
	return apis, nil
}

func GetAllApiCount() int64 {
	var count int64
	err := DB.Table("sys_apis").Count(&count).Error
	if err != nil {
		log.Printf("获取api数量错误:%s", err.Error())
		return 0
	}
	return count
}
