package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my.blogs/service"
	"github.com/my.blogs/utils/myErrors"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	UserApi
	PostApi
	CommentApi
}

var (
	userService    = service.ServiceGroupApp.UserService
	postService    = service.ServiceGroupApp.PostService
	commentService = service.ServiceGroupApp.CommentService
)

func GetUserIdByContext(c *gin.Context) uint {
	userId, ok := c.Get("userId")
	if !ok {
		panic(myErrors.ResponseError{Message: "服务器错误UserId"})
	}
	return userId.(uint)
}
