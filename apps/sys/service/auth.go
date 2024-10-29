package service

import (
	"errors"
	"github.com/ZLinFeng/play-server/apps/sys/model/request"
	"github.com/ZLinFeng/play-server/global"
	"github.com/ZLinFeng/play-server/model/db/system"
	"github.com/ZLinFeng/play-server/utils/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	if common.BcryptHash(userReq.Password) != user.Password {
		return false, errors.New("username or password is wrong")
	}

	return true, nil
}

func (service *AuthService) Logout() bool {
	return true
}

func (service *AuthService) Register(userReq *request.UserReq) (bool, error) {
	if len(userReq.Username) == 0 || len(userReq.Password) == 0 {
		return false, errors.New("username or password is empty")
	}
	var user system.SysUser
	err := global.DB.Where("username = ?", userReq.Username).First(&user).Error
	if err != nil {
		return false, errors.New("user exist already")
	}

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		user = system.SysUser{
			Username: userReq.Username,
			Password: common.BcryptHash(userReq.Password),
			DeptId:   userReq.DeptId,
			Avatar:   userReq.Avatar,
		}

		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		userRoles := make([]system.SysUserRole, 0)
		for _, roleId := range userReq.RoleIds {
			userRoles = append(userRoles, system.SysUserRole{
				UserId: user.ID,
				RoleId: roleId,
			})
		}
		if err := tx.Clauses(clause.Insert{
			Modifier: "IGNORE",
		}).Create(&userRoles).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, errors.New("register user fail")
	}

	return true, nil
}
