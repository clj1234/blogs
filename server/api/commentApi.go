package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my.blogs/model"
	"github.com/my.blogs/model/request"
)

type CommentApi struct {
}

func (api *CommentApi) CreateCommentsByPostId(c *gin.Context) {
	requestComment := new(request.RequestComment)
	userId := GetUserIdByContext(c)
	err := c.ShouldBindJSON(requestComment)
	if err != nil {
		panic(err)
	}
	commentService.CreateCommentsByPostId(&model.Comment{UserId: userId, Content: requestComment.Content, PostId: requestComment.PostId})
	c.JSON(200, gin.H{})
}

func (api *CommentApi) GetCommentsListByPostId(c *gin.Context) {
	requestComment := new(request.RequestComment)
	err := c.ShouldBindJSON(requestComment)
	if err != nil {
		panic(err)
	}
	comments := commentService.GetCommentsListByPostId(requestComment.PostId)
	c.JSON(200, gin.H{"comments": comments})
}
