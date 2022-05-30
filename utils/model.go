package main

type Sys_apis struct {
	Id          int    `:"id" json:"id"`
	CreatedAt   string `:"created_at" json:"created_at"`
	UpdatedAt   string `:"updated_at" json:"updated_at"`
	DeletedAt   string `:"deleted_at" json:"deleted_at"`
	Path        string `:"path" json:"path"`               // api路径
	Description string `:"description" json:"description"` // api中文描述
	ApiGroup    string `:"api_group" json:"api_group"`     // api组
	Method      string `:"method" json:"method"`           // 方法
}

func (*Sys_apis) TableName() string {
	return "sys_apis"
}
