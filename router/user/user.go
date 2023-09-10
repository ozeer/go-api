package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ozeer/go-api/api"
)

type UserRouter struct{}

func (e *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	userRouter := Router.Group("user")
	userApi := api.ApiGroupApp.UserApiGroup.UserApi
	{
		userRouter.GET("analysis", userApi.Analysis)
	}
	return userRouter
}
