package middleware

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/my.blogs/global"
	"github.com/my.blogs/model/system"
	"github.com/my.blogs/utils/myErrors"
	"gorm.io/gorm"
)

func DefaultError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch e := err.(type) {
				case myErrors.ResponseError:
					c.JSON(500, gin.H{"errorCode": e.ErrCode, "error": e.Message})
					err := gorm.G[system.ErrorLog](global.GORM_DB).Create(context.Background(), &system.ErrorLog{Type: 1, ErrContent: e.Message})
					if err != nil {
						debug.PrintStack()
						return
					}
				default:
					c.JSON(500, gin.H{"error": "服务器内部错误"})
					stackStr := string(debug.Stack())
					err := gorm.G[system.ErrorLog](global.GORM_DB).Create(context.Background(), &system.ErrorLog{Type: 2, ErrContent: stackStr})
					if err != nil {
						debug.PrintStack()
						return
					}
					fmt.Println("Panic:", err, "\nStack:", stackStr)
				}
			}
		}()
		c.Next()
	}
}
