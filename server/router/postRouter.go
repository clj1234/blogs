package router

import (
	"github.com/gin-gonic/gin"
)

func PostRouter(privateRouter *gin.RouterGroup) {
	privateRouter.POST("createPost", postApi.CreatePost)
	privateRouter.POST("getPostList", postApi.GetPostList)
	privateRouter.POST("getPostInfoById", postApi.GetPostInfoById)
	privateRouter.POST("updatePostById", postApi.UpdatePostById)
}
