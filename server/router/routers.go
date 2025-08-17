package router

import (
	gin "github.com/gin-gonic/gin"
	"github.com/my.blogs/global"
	"github.com/my.blogs/middleware"
)

func Routers() *gin.Engine {
	router := gin.New()
	router.Use(middleware.DefaultError())
	privateRouter := router.Group("")
	publicRouter := router.Group("")
	privateRouter.Use(middleware.JWTAuthMiddleware())
	UserRouter(privateRouter, publicRouter)
	PostRouter(privateRouter)
	CommentRouter(privateRouter)
	err := router.Run(global.GLOBAL_CONFIG.System.Addr + ":" + global.GLOBAL_CONFIG.System.Port)
	if err != nil {
		panic(err)
	}
	return router
}
