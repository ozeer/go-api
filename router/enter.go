package router

type RouterGroup struct {
	User   UserRouter
	System SystemRouter
}

var RouterGroupApp = new(RouterGroup)
