package service

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	UserService
	PostService
	CommentService
}
