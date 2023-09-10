package router

import (
	"github.com/ozeer/go-api/router/system"
	"github.com/ozeer/go-api/router/user"
)

type RouterGroup struct {
	User   user.RouterGroup
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
