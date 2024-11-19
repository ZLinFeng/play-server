package initial

import (
	"github.com/ZLinFeng/play-server/apps"
	"github.com/ZLinFeng/play-server/global"
	"github.com/ZLinFeng/play-server/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(middleware.FormatRecovery()).Use(middleware.Cors())

	privateRouterGroup := router.Group(global.AppConfig.Server.RoutePrefix)
	publicRouterGroup := router.Group(global.AppConfig.Server.RoutePrefix)

	// 初始化系统router
	apps.RouterGroupApp.System.InitSystemRouter(publicRouterGroup, privateRouterGroup)

	global.Routes = router.Routes()

	global.Logger.Info("init router success")
	return router
}
