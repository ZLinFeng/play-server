package api

import (
	"github.com/ZLinFeng/play-server/model/common/response"
	"github.com/gin-gonic/gin"
)

type AuthApi struct{}

func (api *AuthApi) Login(c *gin.Context) {
	authService.Login("a", "b")
	response.Ok(c)
}
