package api

import (
	"github.com/ZLinFeng/play-server/apps/sys/model/request"
	"github.com/ZLinFeng/play-server/model/common/response"
	"github.com/gin-gonic/gin"
)

type AuthApi struct{}

func (api *AuthApi) Login(c *gin.Context) {
	var userReq request.UserReq

	if err := c.ShouldBindBodyWithJSON(&userReq); err != nil {
		response.FailWithDetailed(response.ErrorEntityCode, response.ErrorEntityMsg, c)
		return
	}

	_, loginErr := authService.Login(&userReq)
	if loginErr != nil {
		response.FailWithDetailed(response.AuthFailCode, response.AuthFailMsg, c)
		return
	}
	response.Ok(c)
}

func (api *AuthApi) Register(c *gin.Context) {
	var userReq request.UserReq
	if err := c.ShouldBindBodyWithJSON(&userReq); err != nil {
		response.FailWithDetailed(response.ErrorEntityCode, response.ErrorEntityMsg, c)
		return
	}
	_, registerErr := authService.Register(&userReq)
	if registerErr != nil {
		response.FailWithDetailed(response.AuthFailCode, response.AuthFailMsg, c)
	}
	response.Ok(c)
}
