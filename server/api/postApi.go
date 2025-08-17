package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my.blogs/model"
	"github.com/my.blogs/model/request"
)

type PostApi struct{}

func (api *PostApi) CreatePost(c *gin.Context) {
	requestPost := new(request.RequestPost)
	userId := GetUserIdByContext(c)
	err := c.ShouldBindJSON(requestPost)
	if err != nil {
		panic(err)
	}
	postService.CreatePost(&model.Post{UserId: userId, Content: requestPost.Content, Title: requestPost.Title})
	c.JSON(200, gin.H{})
}

func (api *PostApi) GetPostList(c *gin.Context) {
	//requestPost := new(request.RequestPost)
	//err := c.ShouldBindJSON(requestPost)
	//if err != nil {
	//	panic(err)
	//}
	posts := postService.GetPostList()
	c.JSON(200, gin.H{"posts": posts})
}

func (api *PostApi) GetPostInfoById(c *gin.Context) {
	requestPost := new(request.RequestPost)
	err := c.ShouldBindJSON(requestPost)
	if err != nil {
		panic(err)
	}
	posts := postService.GetPostInfoById(requestPost.PostId)
	c.JSON(200, gin.H{"posts": posts})
}

func (api *PostApi) UpdatePostById(c *gin.Context) {
	requestPost := new(request.RequestPost)
	userId := GetUserIdByContext(c)
	err := c.ShouldBindJSON(requestPost)
	if err != nil {
		panic(err)
	}
	post := model.Post{Title: requestPost.Title, Content: requestPost.Content}
	post.ID = requestPost.PostId
	postService.UpdatePostById(&post, userId)
	c.JSON(200, gin.H{})
}
