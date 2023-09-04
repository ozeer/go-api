package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ozeer/go-api/api"
)

type AuthRouter struct{}

func (e *UserRouter) InitCourseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	authRouter := Router.Group("auth")
	jwtApi := api.ApiGroupApp.AuthApiGroup.JwtApi
	{
		authRouter.POST("check", jwtApi.Check)
	}
	return authRouter
}
