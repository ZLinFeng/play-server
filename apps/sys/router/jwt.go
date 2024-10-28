package router

import "github.com/gin-gonic/gin"

type JwtRouter struct{}

func (router *JwtRouter) InitJwtRouter(pri *gin.RouterGroup) {
	jwtRouterGroup := pri.Group("/jwt")

	jwtRouterGroup.GET("")
}
