package helper

import (
	"GinVueA/define"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(id uint, name string, expireAt int64) (string, error) {
	uc := define.UserClaim{
		Id:   id,
		Name: name,
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
