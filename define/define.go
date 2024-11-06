package define

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	// Jwtkey 密钥
	Jwtkey = "sys-admin"
	// TokenExpire token有效期，7天
	TokenExpire = time.Now().Add(time.Second * 3600 * 24 * 7).Unix()
	// RefreshTokenExpire 刷新token有效期14天
	RefreshTokenExpire = time.Now().Add(time.Second * 3600 * 24 * 14).Unix()
	// DefaultSize 默认分页没有显示条数
	DefaultSize = 10
	// StaticResource 静态文件目录
	StaticResource = "../../go-admin-static"
	// EmailFrom 邮件发送方
	EmailFrom = ""
	// EmailPassWord 邮箱授权码
	EmailPassWord = ""
	// EmailHost 邮箱host
	EmailHost = "smtp.qq.com"
	// EmailPort 邮箱端口号
	EmailPort = "587"
	// DbPath ip2region存放路径
	DbPath = StaticResource + "/ip2region/ip2region.xdb"
)

type UserClaim struct {
	Id      uint
	Name    string
	IsAdmin bool // 是否超管
	RoleId  uint // 所属角色
	jwt.StandardClaims
}
