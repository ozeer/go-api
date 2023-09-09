package cache

import (
	"context"
	"fmt"

	"github.com/ozeer/go-api/global"
)

var USER_ANALYSIS_AGE = "user_analysis_age"

var RdsData = map[string]string{
	"under_20": "",
	"20-30":    "",
	"30-40":    "",
	"40-50":    "",
	"above_50": "",
}

func SetAnalysisAgeGroup(values map[string]string) {
	err := global.REDIS.HMSet(context.Background(), USER_ANALYSIS_AGE, values).Err()
	if err != nil {
		fmt.Println("写入数据到Redis时发生错误:", err)
		return
	}
}
