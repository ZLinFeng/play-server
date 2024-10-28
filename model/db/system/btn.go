package system

import "github.com/ZLinFeng/play-server/model/common"

type SysMenuBtn struct {
	common.BaseDbModel
	Name   string `json:"name" gorm:"comment:按钮名称"`
	Desc   string `json:"desc" gorm:"comment:按钮备注"`
	MenuId uint64 `json:"menuId" gorm:"comment:所属菜单ID"`
}

func (SysMenuBtn) TableName() string {
	return "sys_menu_btns"
}

type SysRoleBtn struct {
	RoleId       uint64     `json:"roleId" gorm:"comment:角色ID"`
	SysMenuId    uint64     `json:"menuId" gorm:"comment:菜单ID"`
	SysMenuBtnId uint64     `json:"btnId" gorm:"comment:按钮ID"`
	SysMenuBtn   SysMenuBtn `json:"sysMenuBtn"`
}
