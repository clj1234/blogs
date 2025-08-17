package router

import "github.com/my.blogs/api"

type RoutersGroup struct {
}

var (
	userApi    = api.ApiGroupApp.UserApi
	postApi    = api.ApiGroupApp.PostApi
	commentApi = api.ApiGroupApp.CommentApi
)
