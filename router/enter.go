package router

import (
	"github.com/ozeer/go-api/router/designer"
	"github.com/ozeer/go-api/router/example"
	"github.com/ozeer/go-api/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Designer designer.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
