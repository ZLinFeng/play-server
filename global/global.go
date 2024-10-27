package global

import (
	"github.com/ZLinFeng/play-server/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	AppConfig *config.Config
	DB        *gorm.DB
	Logger    *zap.Logger
)
