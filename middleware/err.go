package middleware

import (
	"github.com/ZLinFeng/play-server/global"
	"github.com/ZLinFeng/play-server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func FormatRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.Error("unexpected panic.", zap.Any("err", err))
				response.FailWithMessage("internal server error", c)
			}
		}()
		c.Next()
	}
}
