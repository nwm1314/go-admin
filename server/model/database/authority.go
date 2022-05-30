package database

type Authority struct {
	CreatedAt     string `gorm:"created_at" json:"created_at"`
	UpdatedAt     string `gorm:"updated_at" json:"updated_at"`
	DeletedAt     string `gorm:"deleted_at" json:"deleted_at"`
	AuthorityId   string `gorm:"authority_id" json:"authority_id"`     // 角色ID
	AuthorityName string `gorm:"authority_name" json:"authority_name"` // 角色名
	ParentId      string `gorm:"parent_id" json:"parent_id"`           // 父角色ID
	DefaultRouter string `gorm:"default_router" json:"default_router"` // 默认菜单
}

func (Authority) TableName() string {
	return "sys_authorities"
}

type AuthorityMenu struct {
	BaseMenuId  uint   `gorm:"column:sys_base_menu_id" json:"base_menu_id"`
	AuthorityId string `gorm:"column:sys_authority_authority_id" json:"authority_id"`
}

func (AuthorityMenu) TableName() string {
	return "sys_authority_menus"
}
