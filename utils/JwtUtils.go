package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

type JwtUtils struct {
}

// JwtCustomerClaims 自定义jwt结构体
type JwtCustomerClaims struct {
	ID   int
	Name string
	jwt.RegisteredClaims
}

// GenerateToken 生成 jwt token string
func (util JwtUtils) GenerateToken(id int, name string) (string, error) {
	expiresTime := viper.GetDuration("jwt.expiresTime") // 过期时间
	// token配置信息
	iJwtCustomerClaims := JwtCustomerClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                // token颁发时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * expiresTime)), // 当前时间往后加一段时间
			Subject:   "JwtToken",
		},
	}
	// token （加密算法，token配置信息）
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustomerClaims)
	// 根据密钥获取签名
	return token.SignedString([]byte(viper.GetString("jwt.signKey")))
}

// ParseToken 解析jwt Token, 得到自定义Jwt结构体
func ParseToken(tokenString string) (JwtCustomerClaims, error) {
	iJwtCustomerClaims := JwtCustomerClaims{}
	// 得到自定义的jwt结构体
	token, err := jwt.ParseWithClaims(tokenString, &iJwtCustomerClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt.signKey")), nil
	})

	if err == nil && !token.Valid {
		err = errors.New("InValid Token ")
	}
	return iJwtCustomerClaims, err
}

// IsValidToken 判断token是否有效
func IsValidToken(tokenString string) bool {
	_, err := ParseToken(tokenString)
	if err != nil {
		return false
	}
	return true
}
