package global

import (
	"github.com/ZLinFeng/play-server/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	AppConfig *config.Config
	DB        *gorm.DB
	Logger    *zap.Logger
	Routes    gin.RoutesInfo
)
