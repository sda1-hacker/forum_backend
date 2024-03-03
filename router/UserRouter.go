package router

import (
	"github.com/forum_backend/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(e *gin.Engine) {
	userGroup := e.Group("/user")
	{
		userGroup.GET("/add", controller.UserController{}.AddUser)
	}
}
