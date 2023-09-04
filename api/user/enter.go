package user

import "github.com/ozeer/go-api/service"

type ApiGroup struct {
	UserApi
}

var (
	loginService = service.ServiceGroupApp.UserServiceGroup.LoginService
)
