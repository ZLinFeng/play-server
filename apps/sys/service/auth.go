package service

import (
	"errors"
	"github.com/ZLinFeng/play-server/global"
	"github.com/ZLinFeng/play-server/model/db/system"
	"github.com/ZLinFeng/play-server/utils/common"
)

type AuthService struct {
}

func (service *AuthService) Login(username, password string) (bool, error) {
	if len(username) == 0 || len(password) == 0 {
		return false, errors.New("username or password is empty")
	}

	var user system.SysUser
	err := global.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return false, errors.New("user or password is wrong")
	}

	if common.BcryptHash(password) != user.Password {
		return false, errors.New("username or password is wrong")
	}

	return true, nil
}

func (service *AuthService) Logout() bool {
	return true
}

func (service *AuthService) Register(username, password string) (bool, error) {
	if len(username) == 0 || len(password) == 0 {
		return false, errors.New("username or password is empty")
	}
	return true, nil
}
