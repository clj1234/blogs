package service

import (
	"context"

	"github.com/my.blogs/global"
	"github.com/my.blogs/model"
	"github.com/my.blogs/utils/myErrors"
	"gorm.io/gorm"
)

type CommentService struct{}

func (commentService *CommentService) CreateCommentsByPostId(comment *model.Comment) {
	if comment == nil || comment.Content == "" {
		panic(myErrors.ResponseError{Message: "评论不可为空"})
	}
	err := gorm.G[model.Comment](global.GORM_DB).Create(context.Background(), comment)
	if err != nil {
		panic(err)
	}
}

func (commentService *CommentService) GetCommentsListByPostId(postId uint) []model.Comment {
	comments, err := gorm.G[model.Comment](global.GORM_DB).Where("post_id = ?", postId).Find(context.Background())
	if err != nil {
		panic(err)
	}
	return comments
}
