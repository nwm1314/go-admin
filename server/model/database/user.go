package database

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uuid        uuid.UUID `json:"uuid"`         // 用户UUID
	Username    string    `json:"username"`     // 用户登录名
	Password    string    `json:"password"`     // 用户登录密码
	NickName    string    `json:"nick_name"`    // 用户昵称
	SideMode    string    `json:"side_mode"`    // 用户侧边主题
	HeaderImg   string    `json:"header_img"`   // 用户头像
	BaseColor   string    `json:"base_color"`   // 基础颜色
	ActiveColor string    `json:"active_color"` // 活跃颜色
	AuthorityId string    `json:"authority_id"` // 用户角色ID
	Phone       string    `json:"phone"`        // 用户手机号
	Email       string    `json:"email"`        // 用户邮箱
}

func (User) TableName() string {
	return "sys_users"
}

type UserAuthority struct {
	UserId      uint   `gorm:"column:sys_user_id" json:"user_id"`
	AuthorityId string `gorm:"column:sys_authority_authority_id" json:"authority_id"`
}

func (UserAuthority) TableName() string {
	return "sys_user_authority"
}
