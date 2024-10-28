package system

type SysUserRole struct {
	UserId uint64 `json:"userId" gorm:"primaryKey"`
	RoleId uint64 `json:"roleId" gorm:"primaryKey"`

	User SysUser `json:"user" gorm:"foreignKey:UserId;references:ID;constraint:OnDelete:CASCADE;"`
	Role SysRole `json:"role" gorm:"foreignKey:RoleId;references:ID;constraint:OnDelete:CASCADE;"`
}

func (SysUserRole) TableName() string {
	return "sys_user_roles"
}
