package user

import "github.com/ozeer/go-api/global"

type User struct {
	global.MODEL
	Uid      uint64 `json:"uid" form:"uid" gorm:"uniqueIndex:uniq_idx_uid,comment:用户uid"`
	UserName string `json:"user_name" form:"user_name" gorm:"uniqueIndex:uniq_idx_nick,comment:用户名"`
	Phone    string `json:"phone" form:"phone" gorm:"uniqueIndex:uniq_idx_p,comment:用户手机号"`
	Password string `json:"password" form:"password" gorm:"comment:密码"`
	Email    string `json:"email" form:"email" gorm:"comment:邮箱"`
	Ip       string `json:"ip" form:"ip" gorm:"comment:注册ip"`
}

func (User) TableName() string {
	return "user"
}
