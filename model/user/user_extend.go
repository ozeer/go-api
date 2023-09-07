package user

import "github.com/ozeer/go-api/global"

type UserExtend struct {
	global.MODEL
	Uid      uint64 `json:"uid" form:"uid" gorm:"uniqueIndex:uniq_idx_uid,comment:用户uid"`
	Sex      uint8  `json:"sex" form:"sex" gorm:"comment:性别 0未知 1男 2女"`
	Birthday string `json:"birthday" form:"birthday" gorm:"comment:出生日期"`
}

func (UserExtend) TableName() string {
	return "user_extend"
}
