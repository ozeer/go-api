package system

import (
	"github.com/gin-gonic/gin"
	"github.com/ozeer/go-api/global"
	"github.com/ozeer/go-api/model/common/response"
	"go.uber.org/zap"
)

type SnowFlakeApi struct{}

type SnowFlakeResponse struct {
	Id uint64 `json:"id" gorm:"comment:雪花ID"`
}

// 获取Id
func (u *SnowFlakeApi) GenId(c *gin.Context) {
	id, err := global.SF.NextID()
	if err != nil {
		global.LOG.Error("snowflake err:", zap.Error(err))
		return
	}

	response.OkWithDetailed(SnowFlakeResponse{Id: id}, "获取成功", c)
}
