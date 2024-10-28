package initial

import (
	"fmt"
	"github.com/ZLinFeng/play-server/config"
	"github.com/ZLinFeng/play-server/global"
	"go.uber.org/zap"
)

func InitServer(c *config.Config) {
	router := InitRouter()
	address := fmt.Sprintf(":%d", c.Server.Port)
	err := router.Run(address)
	if err != nil {
		global.Logger.Error("fatal error while start server", zap.Error(err))
	}
}
