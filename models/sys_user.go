package models

// User
type SysUser struct {
	Model
	Username     string  `gorm:"unique;comment:'用户名'" json:"username"`
	Password     string  `gorm:"comment:'密码'" json:"password"`
	Mobile       string  `gorm:"comment:'手机'" json:"mobile"`
	Avatar       string  `gorm:"comment:'头像'" json:"avatar"`
	Nickname     string  `gorm:"comment:'昵称'" json:"nickname"`
	Introduction string  `gorm:"comment:'自我介绍'" json:"introduction"`
	Status       *bool   `gorm:"type:tinyint(1);default:1;comment:'用户状态(正常/禁用, 默认正常)'" json:"status"` // 由于设置了默认值, 这里使用ptr, 可避免赋值失败
	Creator      string  `gorm:"comment:'创建人'" json:"creator"`
	RoleId       uint    `gorm:"comment:'角色Id外键'" json:"roleId"`
	Role         SysRole `gorm:"foreignkey:RoleId" json:"role"` // 将SysUser.RoleId指定为外键
}

func (m SysUser) TableName() string {
	return m.Model.TableName("sys_user")
}
