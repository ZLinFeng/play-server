package system

import "github.com/ZLinFeng/play-server/model/common"

type SysUser struct {
	common.BaseDbModel
	DeptId   uint64         `json:"deptId"`
	Username string         `json:"username" gorm:"index;comment:用户名"`
	Password string         `json:"password" gorm:"comment:密码"`
	Avatar   string         `json:"avatar" gorm:"comment:头像地址"`
	Enable   int            `json:"enable" gorm:"default:1;comment:1表示可用,2表示冻结"`
	Roles    []SysRole      `json:"roles" gorm:"many2many:sys_user_roles;"`
	Dept     SysDept        `json:"dept" gorm:"foreignKey:DeptId;references:ID;"`
	Settings common.JsonMap `json:"settings" gorm:"type:text;default:null;comment:用户配置"`
}

func (SysUser) TableName() string {
	return "sys_users"
}
