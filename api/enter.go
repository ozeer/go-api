package api

import (
	"github.com/ozeer/go-api/api/auth"
	"github.com/ozeer/go-api/api/user"
)

type ApiGroup struct {
	UserApiGroup user.ApiGroup
	AuthApiGroup auth.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
