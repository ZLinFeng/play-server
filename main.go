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

	//配置日志
	global.Logger = initial.InitLog(&global.AppConfig.Log)

	//数据库
	//global.DB = initial.InitDb(&global.AppConfig.SysMysql)
}
