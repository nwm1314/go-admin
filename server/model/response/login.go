package response

type MenuParameter struct {
	CreatedAt     string `json:"createdAt"`
	Id            uint   `json:"id"`
	Key           string `json:"key"`
	SysBaseMenuID int    `json:"sysBaseMenuID"`
	Type          string `json:"type"`
	UpdatedAt     string `json:"updatedAt"`
	Value         string `json:"value"`
}
type MenuBtb struct {
	CreatedAt     string `json:"createdAt"`
	Desc          string `json:"desc"`
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	SysBaseMenuID int    `json:"sysBaseMenuID"`
	UpdatedAt     string `json:"updatedAt"`
}

type Btn struct {
}

type Meta struct {
	KeepAlive   bool   `json:"keepAlive"`
	DefaultMenu bool   `json:"defaultMenu"`
	Title       string `json:"title"`
	Icon        string `json:"icon"`
	CloseTab    bool   `json:"closeTab"`
}
type Menu struct {
	Authorities []Authority     `json:"authoritys"`
	Children    []Menu          `json:"children"`
	Component   string          `json:"component"`
	CreatedAt   string          `json:"CreatedAt"`
	Hidden      bool            `json:"hidden"`
	Id          uint            `json:"ID"`
	Meta        Meta            `json:"meta"`
	MenuBtn     []MenuBtb       `json:"menuBtn"`
	Name        string          `json:"name"`
	Parameters  []MenuParameter `json:"Parameters"`
	Btns        []Btn           `json:"btns"`
	ParentId    string          `json:"parentId"`
	Path        string          `json:"path"`
	Sort        int             `json:"sort"`
	UpdatedAt   string          `json:"UpdatedAt"`
	MenuId      string          `json:"menuId"`
}

type MenuResp struct {
	Menus []Menu `json:"menus"`
}
type Authority struct {
	AuthorityId     string      `json:"authorityId"`
	AuthorityName   string      `json:"authorityName"`
	Children        []Authority `json:"children"`
	CreatedAt       string      `json:"CreatedAt"`
	DataAuthorityId []Authority `json:"dataAuthorityId"`
	DefaultRouter   string      `json:"defaultRouter"`
	DeletedAt       string      `json:"DeletedAt"`
	Menus           []Menu      `json:"menus"`
	ParentId        string      `json:"parentId"`
	UpdatedAt       string      `json:"UpdatedAt"`
}

//type User struct {
//	ID          uint      `gorm:"primarykey" json:"id"`
//	CreatedAt   time.Time `json:"createdAt"`
//	UpdateAt    time.Time `json:"updateAt"`
//	Uuid        uuid.UUID `json:"uuid"`         // 用户UUID
//	Username    string    `json:"userName"`     // 用户登录名
//	NickName    string    `json:"nick_name"`    // 用户昵称
//	SideMode    string    `json:"side_mode"`    // 用户侧边主题
//	HeaderImg   string    `json:"header_img"`   // 用户头像
//	BaseColor   string    `json:"base_color"`   // 基础颜色
//	ActiveColor string    `json:"active_color"` // 活跃颜色
//	AuthorityId string    `json:"authority_id"` // 用户角色ID
//	Phone       string    `json:"phone"`        // 用户手机号
//	Email       string    `json:"email"`        // 用户邮箱
//}

type LoginResp struct {
	ExpiresAt int64    `json:"expiresAt"`
	Token     string   `json:"token"`
	User      UserInfo `json:"user"`
}
