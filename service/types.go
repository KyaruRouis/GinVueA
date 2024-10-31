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
