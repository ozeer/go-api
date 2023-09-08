package utils

import (
	"math/rand"
	"time"
)

const (
	// 随机日期格式
	DATE_TEMPLATE = "2006-01-02"
)

// 生成随机日期
// 参数：起始年份，结束年份
func GenerateRandomDate(startYear, endYear int) string {
	// 设置随机种子，以便每次运行生成不同的随机日期
	// rand.Seed(time.Now().UnixNano())
	randNew := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 生成随机年份
	year := randNew.Intn(endYear-startYear+1) + startYear

	// 生成随机月份
	month := randNew.Intn(12) + 1

	// 生成随机日期
	daysInMonth := 31 // 默认31天
	switch month {
	case 4, 6, 9, 11:
		daysInMonth = 30
	case 2:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			daysInMonth = 29 // 闰年2月29天
		} else {
			daysInMonth = 28 // 平年2月28天
		}
	}

	day := randNew.Intn(daysInMonth) + 1

	// 创建时间对象
	randomDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	return randomDate.Format(DATE_TEMPLATE)
}

// 生成随机整数
func GenRandomNumber(n int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(n)
}
