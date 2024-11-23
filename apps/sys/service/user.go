package service

import (
	"errors"

	"github.com/ZLinFeng/play-server/apps/sys/model/request"
	"github.com/ZLinFeng/play-server/global"
	"github.com/ZLinFeng/play-server/model/db/system"
	"github.com/ZLinFeng/play-server/utils/common"
	"gorm.io/gorm"
)

type UserService struct{}

func (s *UserService) AddUser(userReq *request.UserReq) (uint64, error) {
	if len(userReq.Username) == 0 || len(userReq.Password) == 0 {
		return 0, errors.New("username or password is empty")
	}
	if userReq.DeptId == 0 || userReq.RoleIds == nil {
		return 0, errors.New("auth or dept is empty")
	}
	var user system.SysUser
	err := global.DB.Where("username = ?", userReq.Username).First(&user).Error
	if err == nil {
		return 0, errors.New("user exist already")
	}

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		user = system.SysUser{
			Username: userReq.Username,
			Password: common.BcryptHash(userReq.Password),
			DeptId:   userReq.DeptId,
			Avatar:   userReq.Avatar,
			Enable:   userReq.Enable,
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
		if err := tx.Create(&userRoles).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, errors.New("add user fail")
	}

	return user.ID, nil
}

func (s *UserService) DeleteUser(id uint64) (bool, error) {
	err := global.DB.Where("id = ?", id).First(&system.SysUser{}).Error
	if err != nil {
		return false, errors.New("user not exist")
	}
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("sys_user_id = ?", id).Delete(&system.SysUserRole{}).Error; err != nil {
			return err
		}
		if err := tx.Where("id = ?", id).Delete(&system.SysUser{}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, errors.New("delete user fail")
	}
	return true, nil
}

func (s *UserService) GetUserById(id int) (string, error) {
	return "user1", nil
}

func (s *UserService) GetUserList() ([]system.SysUser, error) {
	return nil, nil
}
