package request

// UserReq 登录、修改密码、注册的请求结构
type UserReq struct {
	Username string   `json:"username" binding:"required"`
	Password string   `json:"password" binding:"required"`
	DeptId   uint64   `json:"deptId"`
	RoleIds  []uint64 `json:"roleIds"`
	Avatar   string   `json:"avatar"`
	Enable   int      `json:"active" binding:"required"`
}
