package router

import (
	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

func (router *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	authRouter := Router.Group("auth")
	authRouter.POST("login", authApi.Login)
	authRouter.POST("register", authApi.Register)
}
