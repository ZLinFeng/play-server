package main

import (
	"github.com/ZLinFeng/play-server/config"
	"github.com/ZLinFeng/play-server/global"
	"github.com/ZLinFeng/play-server/initial"
)

func main() {
	// 初始化配置
	global.AppConfig = config.InitConfig()

	initial.PrintBanner(global.AppConfig)

	// 初始化日志
	global.Logger = initial.InitLog(&global.AppConfig.Log)

	// 初始化数据库
	global.DB = initial.InitDb(&global.AppConfig.SysMysql)
	initial.RegisterTables()

	// 启动服务
	initial.InitServer(global.AppConfig)
}
