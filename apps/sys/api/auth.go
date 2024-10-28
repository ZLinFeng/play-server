package api

import (
	"github.com/ZLinFeng/play-server/model/common/response"
	"github.com/gin-gonic/gin"
)

type AuthApi struct{}

func (api *AuthApi) Login(username, password string, c *gin.Context) {
	response.Ok(c)
}
