package api

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/my.blogs/model"
	"github.com/my.blogs/model/request"
)

type UserApi struct{}

func (userApi *UserApi) CreateUser(c *gin.Context) {
	userRequest := new(request.UserRequest)

	err := c.ShouldBindJSON(userRequest)
	if err != nil {
		panic(err)
	}
	fmt.Println("userRequest:", userRequest)
	user, err := userService.CreateUser(&model.User{Username: userRequest.Username, Password: userRequest.Password})
	if err != nil {
		c.String(500, err.Error())
	}
	s, _ := json.Marshal(user)
	c.String(200, string(s))
}

func (userApi *UserApi) Login(c *gin.Context) {
	userRequest := new(request.UserRequest)
	err := c.ShouldBindJSON(userRequest)
	if err != nil {
		panic(err)
	}
	tokenString, err := userService.Login(&model.User{Username: userRequest.Username, Password: userRequest.Password})
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{"U-TOKEN": tokenString})
}
