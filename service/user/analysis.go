package user

import (
	"context"

	"github.com/ozeer/go-api/cache"
	"github.com/ozeer/go-api/global"
	"go.uber.org/zap"
)

type AnalysisService struct{}

var AgeGroups = map[string]int{
	"under_20": 0,
	"20-30":    0,
	"30-40":    0,
	"40-50":    0,
	"above_50": 0,
}

func (a *AnalysisService) GetAnalysisUserByAge() map[string]string {
	data, err := global.REDIS.HGetAll(context.Background(), cache.USER_ANALYSIS_AGE).Result()

	if err != nil {
		global.LOG.Error("获取用户分析数据失败", zap.Error(err))
	}

	return data
}
