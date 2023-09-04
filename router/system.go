package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ozeer/go-api/api"
)

type SystemRouter struct{}

func (e *SystemRouter) InitSystemRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	systemRouter := Router.Group("base")
	jwtApi := api.ApiGroupApp.AuthApiGroup.JwtApi
	{
		systemRouter.POST("check", jwtApi.Check)
	}
	return systemRouter
}
