package response

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type UserInfoResp struct {
	UserInfo UserInfo `json:"userInfo"`
}

type UserInfo struct {
	ID          uint        `gorm:"primarykey" json:"ID"`
	CreatedAt   time.Time   `json:"CreatedAt"`
	UpdateAt    time.Time   `json:"UpdateAt"`
	Uuid        uuid.UUID   `json:"uuid"`         // 用户UUID
	Username    string      `json:"userName"`     // 用户登录名
	NickName    string      `json:"nick_name"`    // 用户昵称
	SideMode    string      `json:"side_mode"`    // 用户侧边主题
	HeaderImg   string      `json:"header_img"`   // 用户头像
	BaseColor   string      `json:"base_color"`   // 基础颜色
	ActiveColor string      `json:"active_color"` // 活跃颜色
	AuthorityId string      `json:"authority_id"` // 用户角色ID
	Phone       string      `json:"phone"`        // 用户手机号
	Email       string      `json:"email"`        // 用户邮箱
	Authorities []Authority `json:"authorities"`
	Authority   Authority   `json:"authority"`
}
