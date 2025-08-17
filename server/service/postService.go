package service

import (
	"context"

	"github.com/my.blogs/global"
	"github.com/my.blogs/model"
	"github.com/my.blogs/utils/myErrors"
	"gorm.io/gorm"
)

type PostService struct{}

func (postService *PostService) CreatePost(post *model.Post) {
	err := gorm.G[model.Post](global.GORM_DB).Create(context.Background(), post)
	if err != nil {
		panic(err)
	}
}

func (postService *PostService) GetPostList() []model.Post {
	var posts []model.Post
	posts, err := gorm.G[model.Post](global.GORM_DB).Find(context.Background())
	if err != nil {
		panic(err)
	}
	return posts
}

func (postService *PostService) GetPostInfoById(postId uint) model.Post {
	//post, err := gorm.G[model.Post](global.GORM_DB).Preload("Comments", func(db gorm.PreloadBuilder) error {
	//	return nil
	//}).Where("id = ?", postId).First(context.Background())
	post, err := gorm.G[model.Post](global.GORM_DB).Where("id = ?", postId).First(context.Background())
	if err != nil {
		panic(err)
	}
	return post
}

func (postService *PostService) UpdatePostById(newPost *model.Post, userId uint) {
	oldPost, err := gorm.G[model.Post](global.GORM_DB).Where("id = ?", newPost.ID).First(context.Background())
	if err != nil {
		panic(err)
	}
	if oldPost.UserId != userId {
		panic(myErrors.ResponseError{Message: "非本用户创建的文章"})
	}
	_, err = gorm.G[model.Post](global.GORM_DB).Where("id = ?", newPost.ID).Updates(context.Background(), *newPost)
	if err != nil {
		panic(err)
	}
}
