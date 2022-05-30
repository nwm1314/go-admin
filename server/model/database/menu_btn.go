package database

import "gorm.io/gorm"

type MenuBtn struct {
	gorm.Model
	Name          string `json:"name"` // 按钮关键key
	Desc          string `json:"desc"`
	SysBaseMenuId int    `json:"sys_base_menu_id"` // 菜单ID
}

func (MenuBtn) TableName() string {
	return "sys_base_menu_btns"
}
