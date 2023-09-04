package auth

import "github.com/ozeer/go-api/service"

type ApiGroup struct {
	JwtApi
}

var (
	loginService = service.ServiceGroupApp.UserServiceGroup.LoginService
)
