package define

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	// 密钥
	Jwtkey = "sys-admin"
	// token有效期，7天
	TokenExpire = time.Now().Add(time.Hour * 24 * 7).Unix()
	// 刷新token有效期，14天
	RefreshTokenExpire = time.Now().Add(time.Hour * 24 * 14).Unix()
	// 默认分页显示条数
	DefaultSize = 10
)

type UserClaim struct {
	Id      uint
	Name    string
	IsAdmin bool // 是否管理员
	jwt.StandardClaims
}
