package service

// 接收登录参数结构体
type LoginPasswordRequest struct {
	UserName string `json:"username"` // 登录用户名
	Password string `json:"password"` // 登录密码
}

// LoginPasswordReply登录成功后的token结构体
type LoginPasswordReply struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

// GetUserListRequest 获取管理员列表参数结构体
type GetUserListRequest struct {
	*QueryRequest
}

// QueryRequest 关键字和分页信息结构体
type QueryRequest struct {
	Page    int    `json:"pageIndex" form:"pageIndex"`
	Size    int    `json:"pageSize" form:"pageSize"`
	Keyword string `json:"keyword" form:"keyword"`
}

// GetUserListReply 返回管理员信息结构体
type GetUserListReply struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// AddUserRequest 接收添加管理员表单数据的结构体
type AddUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Remarks  string `json:"remark"`
}

// GetUserDetailReply 获取管理员信息结构体
type GetUserDetailReply struct {
	ID uint `json:"id"`
	AddUserRequest
}

// UpdateUserRequest 接收更新管理员信息参数的结构体
type UpdateUserRequest struct {
	ID uint `json:"id"`
	AddUserRequest
}

// GetRoleListRequest 获取角色列表查询参数结构体
type GetRoleListRequest struct {
	*QueryRequest
}

// GetRoleListReply 返回角色列表数据结构体
type GetRoleListReply struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Sort      int    `json:"sort"`
	IsAdmin   int    `json:"is_admin"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// AddRoleRequest 新增角色参数结构体
type AddRoleRequest struct {
	Name    string `json:"name"`
	Sort    int64  `json:"sort"`
	IsAdmin int8   `json:"isAdmin"`
	Remarks string `json:"remarks"`
}

// GetRoleDetailReply 返回角色详情信息结构体
type GetRoleDetailReply struct {
	ID uint `json:"id"`
	AddRoleRequest
}

// UpdateRoleRequest 更新角色信息结构体
type UpdateRoleRequest struct {
	ID uint `json:"id"`
	AddRoleRequest
}
