package request

type RegisterInfo struct {
	UserName string `json:"username" form:"username" gorm:"comment:用户名"`
	Phone    string `json:"phone" form:"phone" gorm:"comment:用户手机号"`
	Password string `json:"password" form:"password" gorm:"comment:密码"`
	Email    string `json:"email" form:"email" gorm:"comment:邮箱"`
	Sex      uint8  `json:"sex" form:"sex" gorm:"comment:性别 0未知 1男 2女"`
	Birthday string `json:"birthday" form:"birthday" gorm:"comment:出生日期"`
}
