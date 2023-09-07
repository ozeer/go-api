package service

import (
	"github.com/ozeer/go-api/service/system"
	"github.com/ozeer/go-api/service/user"
)

type ServiceGroup struct {
	UserServiceGroup   user.ServiceGroup
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
