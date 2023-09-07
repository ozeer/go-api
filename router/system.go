package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ozeer/go-api/api"
)

type SystemRouter struct{}

func (e *SystemRouter) InitSystemRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	systemRouter := Router.Group("system")
	systemApi := api.ApiGroupApp.SystemApiGroup.SnowFlakeApi
	{
		systemRouter.GET("genId", systemApi.GenId)
	}
	return systemRouter
}
