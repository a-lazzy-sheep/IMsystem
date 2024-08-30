package utils

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

// 定义一个密钥，用于签名JWT
var jwtKey = []byte("123wyy321")

// Claims 结构体用于存储JWT的声明
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// GenerateToken 生成JWT令牌
func GenerateToken(email, password string) (string, error) {
	// 设置令牌过期时间
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: email + password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用密钥签名令牌并获取完整的编码后的字符串令牌
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
