package helper

import (
	"GinVueA/define"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

// GenerateToken 生成token
func GenerateToken(id uint, roleId uint, name string, expireAt int64) (string, error) {
	uc := define.UserClaim{
		Id:     id,
		Name:   name,
		RoleId: roleId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.Jwtkey))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

// AnalyzeToken 解析token
func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.Jwtkey), nil
	})

	if err != nil {
		return nil, err
	}

	if !claims.Valid {
		return uc, errors.New("token不正确")
	}

	return uc, err

}
