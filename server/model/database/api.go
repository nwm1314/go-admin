package database

type Api struct {
	Id          int    `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
	Path        string `json:"path"`        // api路径
	Description string `json:"description"` // api中文描述
	ApiGroup    string `json:"api_group"`   // api组
	Method      string `json:"method"`      // 方法
}

func (*Api) TableName() string {
	return "sys_apis"
}
