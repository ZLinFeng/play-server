package router

import "github.com/gin-gonic/gin"

type AuthRouter struct{}

func (router *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {}
