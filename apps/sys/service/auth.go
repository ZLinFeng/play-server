package service

import (
	"errors"

	"github.com/ZLinFeng/play-server/apps/sys/model/request"
	"github.com/ZLinFeng/play-server/global"
	"github.com/ZLinFeng/play-server/model/db/system"
	"github.com/ZLinFeng/play-server/utils/common"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
}

func (service *AuthService) Login(userReq *request.UserReq) (bool, error) {
	if len(userReq.Username) == 0 || len(userReq.Password) == 0 {
		return false, errors.New("username or password is empty")
	}

	var user system.SysUser
	err := global.DB.Where("username = ?", userReq.Username).First(&user).Error
	if err != nil {
		return false, errors.New("user or password is wrong")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password))
	if err != nil {
		return false, errors.New("username or password is wrong")
	}

	return true, nil
}

func (service *AuthService) Logout() bool {
	return true
}

func (service *AuthService) ChangePassword(userReq *request.UserReq) (bool, error) {
	if len(userReq.Username) == 0 || len(userReq.Password) == 0 {
		return false, errors.New("username or password is empty")
	}
	password := common.BcryptHash(userReq.Password)
	err := global.DB.Model(&system.SysUser{}).Where(
		"username = ?", userReq.Username).Update(
		"password", password).Error
	if err != nil {
		return false, errors.New("change password fail")
	}
	return true, nil
}

func (service *AuthService) Register(userReq *request.UserReq) (bool, error) {
	return false, nil
}
