package system

import (
	"github.com/gin-gonic/gin"
	"github.com/ozeer/go-api/api"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := api.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.GET("genId", baseApi.GenId)
		baseRouter.POST("register", baseApi.Register)
	}

	return baseRouter
}
