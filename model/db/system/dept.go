package system

import "github.com/ZLinFeng/play-server/model/common"

type SysDept struct {
	common.BaseDbModel
	Name     string    `json:"name" gorm:"comment:部门名"`
	Level    uint      `json:"level" gorm:"comment:部门等级"`
	Sort     uint      `json:"sort" gorm:"comment:排序依据"`
	ParentId uint64    `json:"parentId" gorm:"comment:父部门ID;default:1"`
	Children []SysDept `json:"children" gorm:"-"`
	Users    []SysUser `json:"users" gorm:"foreignKey:DeptId"`
}

func (SysDept) TableName() string {
	return "sys_depts"
}
