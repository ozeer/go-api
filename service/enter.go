package service

import (
	"github.com/ozeer/go-api/service/auth"
	"github.com/ozeer/go-api/service/user"
)

type ServiceGroup struct {
	UserServiceGroup user.ServiceGroup
	AuthServiceGroup auth.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
