package system

import "github.com/ZLinFeng/play-server/model/common"

type SysRole struct {
	common.BaseDbModel
	RoleName string    `json:"roleName" gorm:"index;comment:角色名"`
	Users    []SysUser `json:"users" gorm:"many2many:sys_user_roles;"`
}

func (SysRole) TableName() string {
	return "sys_roles"
}
