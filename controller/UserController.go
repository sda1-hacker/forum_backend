package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (c UserController) AddUser(context *gin.Context) {

	// 这里是service
	// service.UserService{}.AddUser()

	context.JSON(200, gin.H{
		"message": "访问成功..",
	})
}
