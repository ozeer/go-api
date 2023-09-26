package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/ozeer/go-api/api"
)

type DemoRouter struct{}

func (e *DemoRouter) InitDemoRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	demoRouter := Router.Group("demo")
	demoApi := api.ApiGroupApp.DemoApiGroup.DemoApi
	{
		demoRouter.GET("multi_request", demoApi.MultiRequest)
	}
	return demoRouter
}
