package service

type SysServiceGroup struct {
	AuthService
	UserService
}

var SysServiceGroupApp = new(SysServiceGroup)
