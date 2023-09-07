package api

import (
	"github.com/ozeer/go-api/api/system"
	"github.com/ozeer/go-api/api/user"
)

type ApiGroup struct {
	UserApiGroup   user.ApiGroup
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
