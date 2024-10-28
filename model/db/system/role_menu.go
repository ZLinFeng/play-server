package system

type SysRoleMenu struct {
	RoleId uint64 `json:"roleId" gorm:"primaryKey;column:sys_role_id"`
	MenuId uint64 `json:"menuId" gorm:"primaryKey;column:sys_menu_id"`

	Role SysRole `json:"role" gorm:"foreignKey:RoleId;references:ID;constraint:OnDelete:CASCADE;"`
	Menu SysMenu `json:"menu" gorm:"foreignKey:MenuId;references:ID;constraint:OnDelete:CASCADE;"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menus"
}
