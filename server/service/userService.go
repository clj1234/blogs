package service

import (
	"context"
	"sync"

	"github.com/my.blogs/global"
	"github.com/my.blogs/model"
	"github.com/my.blogs/utils"
	"github.com/my.blogs/utils/myErrors"
	"gorm.io/gorm"
)

type UserService struct{}

// TokenUserIdMap map<token,userId>
var TokenUserIdMap = make(map[string]uint, 10)
var tokenUserIdMapLock sync.RWMutex

// GetUserIdByToken 根据token获取用户id /**
func (userService *UserService) GetUserIdByToken(token string) uint {
	tokenUserIdMapLock.RLock()
	defer tokenUserIdMapLock.RUnlock()
	userId := TokenUserIdMap[token]
	return userId
}

// CreateUser 创建用户
func (userService *UserService) CreateUser(user *model.User) (*model.User, error) {
	if user.Password == "" || user.Username == "" {
		panic(myErrors.ResponseError{Message: "用户名、密码不能为空"})
	}
	_, err := gorm.G[model.User](global.GORM_DB).Where("username", user.Username).First(context.Background())
	if err == nil {
		panic(myErrors.ResponseError{Message: "用户名已存在"})
	}
	user.Password = utils.BcryptHash(user.Password)
	err = gorm.G[model.User](global.GORM_DB).Create(context.Background(), user)
	return user, err
}

func (userService *UserService) Login(user *model.User) (tokenString string, err error) {
	newUser, err := gorm.G[model.User](global.GORM_DB).Where("username", user.Username).First(context.Background())
	if err != nil {
		panic(err)
	}
	ok := utils.BcryptCheck(user.Password, newUser.Password)
	if !ok {
		panic(myErrors.NewResponseError("用户名或密码错误"))
	}
	token, err := utils.GenerateToken(newUser.ID)
	if err != nil {
		panic(err)
	}
	return token, nil
}
