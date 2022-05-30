package database

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	MenuLevel   int    `json:"menu_level"`
	ParentId    string `json:"parent_id"`    // 父菜单ID
	Path        string `json:"path"`         // 路由path
	Name        string `json:"name"`         // 路由name
	Hidden      int    `json:"hidden"`       // 是否在列表隐藏
	Component   string `json:"component"`    // 对应前端文件路径
	Sort        int    `json:"sort"`         // 排序标记
	KeepAlive   int    `json:"keep_alive"`   // 附加属性
	DefaultMenu int    `json:"default_menu"` // 附加属性
	Title       string `json:"title"`        // 附加属性
	Icon        string `json:"icon"`         // 附加属性
	CloseTab    int    `json:"close_tab"`    // 附加属性
}

func (Menu) TableName() string {
	return "sys_base_menus"
}
