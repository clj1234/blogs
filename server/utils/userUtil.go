package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/my.blogs/global"
	"golang.org/x/crypto/bcrypt"
)

type CustomClaims struct {
	UserID               uint `json:"user_id"`
	jwt.RegisteredClaims      // 嵌入标准声明（v5版本特性）
}

func GenerateToken(userID uint) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			Issuer:    "go_learn_blogs",                                   // 签发者
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.GLOBAL_CONFIG.System.SigningKey))
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(global.GLOBAL_CONFIG.System.SigningKey), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
