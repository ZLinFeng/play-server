package sys

import (
	"github.com/ZLinFeng/play-server/apps/sys/api"
	"github.com/ZLinFeng/play-server/apps/sys/router"
	"github.com/gin-gonic/gin"
)

type SystemRouterGroup struct {
	router.AuthRouter
}

type SystemApiGroup struct {
	api.AuthApi
}

var SystemApiGroupApp = new(SystemApiGroup)

var (
	authApi = SystemApiGroupApp.AuthApi
)

func (g *SystemRouterGroup) InitRouter(pub *gin.RouterGroup, pri *gin.RouterGroup) {
	g.InitAuthRouter(pub)
}
