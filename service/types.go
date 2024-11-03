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
	UserName string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Remark   string `json:"remark"`
}
