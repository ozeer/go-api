package system

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ozeer/go-api/global"
	"github.com/ozeer/go-api/model/common/response"
	"go.uber.org/zap"
)

type SnowFlakeResponse struct {
	Id uint64 `json:"id" gorm:"comment:雪花ID"`
}

type FakePhoneResponse struct {
	Phone string `json:"phone"`
}

// 获取Id
func (u *BaseApi) GenId(c *gin.Context) {
	id, err := global.SF.NextID()
	if err != nil {
		global.LOG.Error("snowflake err:", zap.Error(err))
		return
	}

	response.OkWithDetailed(SnowFlakeResponse{Id: id}, "获取成功", c)
}

// 获取fake手机号
func (u *BaseApi) GenFakePhone(c *gin.Context) {
	phone, err := global.REDIS.SPop(context.Background(), "phone_pool").Result()

	if err != nil {
		log.Println("错误:", err)
	}
	response.OkWithDetailed(FakePhoneResponse{Phone: phone}, "获取成功", c)
}
