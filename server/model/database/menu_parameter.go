package database

import "gorm.io/gorm"

type MenuParameter struct {
	gorm.Model
	SysBaseMenuId int    `json:"sys_base_menu_id"`
	Type          string `json:"type"`  // 地址栏携带参数为params还是query
	Key           string `json:"key"`   // 地址栏携带参数的key
	Value         string `json:"value"` // 地址栏携带参数的值
}

func (MenuParameter) TableName() string {
	return "sys_base_menu_parameters"
}
