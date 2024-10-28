package api

import "github.com/ZLinFeng/play-server/apps/sys/service"

type SysApiGroup struct {
	AuthApi
	JwtApi
}

var (
	authService = &service.SysServiceGroupApp.AuthService
)

var SysApiGroupApp = new(SysApiGroup)
