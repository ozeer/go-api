package global

import (
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/gorm"
)

const (
	// 数据状态(0:删除 1:正常)
	STATUS_NORMAL = 1
	STATUS_DELETE = 0

	// 日期模版
	DATE_TEMPLATE = "2006-01-02 15:04:05"
)

type GVA_MODEL struct {
	ID        uint           `json:"id" gorm:"primarykey"` // 主键ID
	CreatedAt LocalTime      `json:"created_at"`           // 创建时间
	UpdatedAt LocalTime      `json:"updated_at"`           // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`       // 删除时间
}

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(DATE_TEMPLATE))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	// 判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
