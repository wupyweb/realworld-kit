package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	jwt.RegisteredClaims
	// 添加额外的字段
	// 例如：用户ID
	UserID   int    `json:"userid"`
	UserName string `json:"username"`
}

// jwt算法密钥
var tokenKey = []byte("your_secret_key")

// 生成jwt token
func GenerateToken(id int, username string) (string, error) {
	payload := Payload{
		UserID: id,
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), // 设置token过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                         // token签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                         // 设置token开始生效时间
		},
	}
	// 设置签名算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	// ..
	return token.SignedString(tokenKey)
}

// 解析token
func ParseToken(tokenstring string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(tokenstring, &Payload{}, func(t *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Payload); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
