package define

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	// Jwtkey 密钥
	Jwtkey = "sys-admin"
	// TokenExpire token有效期，7天
	TokenExpire = time.Now().Add(time.Hour * 24 * 7).Unix()
	// RefreshTokenExpire 刷新token有效期，14天
	RefreshTokenExpire = time.Now().Add(time.Hour * 24 * 14).Unix()
	// DefaultSize 默认分页显示条数
	DefaultSize = 10
)

type UserClaim struct {
	Id      uint
	Name    string
	IsAdmin bool // 是否超管
	RoleId  uint // 所属角色
	jwt.StandardClaims
}
