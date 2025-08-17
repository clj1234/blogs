package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserRouter(privateRouter *gin.RouterGroup, publicRouter *gin.RouterGroup) {
	publicRouter.GET("test", func(c *gin.Context) {
		fmt.Println("success")
		c.JSON(200, gin.H{"hello": "world2"})
	})
	publicRouter.POST("createUser", userApi.CreateUser)
	publicRouter.POST("login", userApi.Login)

}
