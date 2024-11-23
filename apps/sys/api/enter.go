package api

import "github.com/ZLinFeng/play-server/apps/sys/service"

type SysApiGroup struct {
	AuthApi
	JwtApi
	UserApi
}

var (
	authService = &service.SysServiceGroupApp.AuthService
	userService = &service.SysServiceGroupApp.UserService
)

var SysApiGroupApp = new(SysApiGroup)
