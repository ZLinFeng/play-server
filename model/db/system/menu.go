package system

import "github.com/ZLinFeng/play-server/model/common"

type SysMenu struct {
	common.BaseDbModel
	Level     uint           `json:"level" gorm:"comment:菜单等级"`
	Sort      uint           `json:"sort" gorm:"comment:排序依据"`
	ParentId  uint           `json:"parentId" gorm:"comment:父菜单ID"`
	Path      string         `json:"path" gorm:"comment:路由"`
	Name      string         `json:"name" gorm:"comment:路由名"`
	Component string         `json:"component" gorm:"前端文件路径"`
	Icon      string         `json:"icon" gorm:"comment:菜单图标"`
	Children  []SysMenu      `json:"children" gorm:"-"`
	SysRoles  []SysRole      `json:"roles" gorm:"many2many:sys_role_menus;"`
	Meta      common.JsonMap `json:"meta" gorm:"type:text;default:null;comment:附加信息"`
}
