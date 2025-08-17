package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/my.blogs/utils"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, _ := c.Cookie("U_TOKEN")
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "未登录"})
			return
		}
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}
		c.Set("userId", claims.UserID) // 存储用户信息到上下文
		c.Next()
	}
}
