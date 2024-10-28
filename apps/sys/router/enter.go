package router

import (
	"github.com/ZLinFeng/play-server/apps/sys/api"
	"github.com/gin-gonic/gin"
)

type SystemRouterGroup struct {
	AuthRouter
	JwtRouter
}

var (
	authApi = &api.SysApiGroupApp.AuthApi
	jwtApi  = &api.SysApiGroupApp.JwtApi
)

func (g *SystemRouterGroup) InitSystemRouter(pub *gin.RouterGroup, pri *gin.RouterGroup) {
	pubSysRouter := pub.Group("sys")
	priSysRouter := pri.Group("sys")
	g.InitAuthRouter(pubSysRouter) // 权限路由
	g.InitJwtRouter(priSysRouter)  // JWT路由
}
