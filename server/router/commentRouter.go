package router

import (
	"github.com/gin-gonic/gin"
)

func CommentRouter(privateRouter *gin.RouterGroup) {
	privateRouter.POST("createCommentsByPostId", commentApi.CreateCommentsByPostId)
	privateRouter.POST("getCommentsListByPostId", commentApi.GetCommentsListByPostId)
}
