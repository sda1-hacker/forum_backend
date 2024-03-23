package router

import (
	"github.com/forum_backend/controller"
	"github.com/forum_backend/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(e *gin.Engine) {
	userGroup := e.Group("/user", middleware.Auth) // 添加了认证
	{
		userGroup.GET("/add", controller.UserController{}.AddUser)
		userGroup.GET("/getUserById/:id", controller.UserController{}.GetUserByID)
	}
}
