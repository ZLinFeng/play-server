package api

import (
	"github.com/ZLinFeng/play-server/apps/sys/model/request"
	"github.com/ZLinFeng/play-server/model/common/response"
	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (api *UserApi) Add(c *gin.Context) {
	var user request.UserReq

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		response.FailWithDetailed(response.ErrorEntityCode, response.ErrorEntityMsg, c)
		return
	}

	userId, err := userService.AddUser(&user)
	if err != nil {
		response.FailWithDetailed(response.AddFailCode, response.AddFailMsg, c)
	} else {
		response.OkWithData(userId, c)
	}
}
