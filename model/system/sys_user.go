package system

import (
	"github.com/gofrs/uuid/v5"
	"github.com/ozeer/go-api/global"
)

type SysUser struct {
	global.MODEL
	UUID     uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`                // 用户UUID
	Username string    `json:"username" gorm:"index;comment:用户登录名"`             // 用户登录名
	Password string    `json:"-"  gorm:"comment:用户登录密码"`                        // 用户登录密码                                      // 活跃颜色
	Phone    string    `json:"phone"  gorm:"comment:用户手机号"`                     // 用户手机号
	Email    string    `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
	Enable   int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
}

func (SysUser) TableName() string {
	return "sys_user"
}
