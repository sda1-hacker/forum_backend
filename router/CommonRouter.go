package router

import (
	"github.com/forum_backend/controller"
	"github.com/gin-gonic/gin"
)

func CommonRouter(e *gin.Engine) {
	common := e.Group("/")
	{
		common.POST("login", controller.UserController{}.Login)
	}
}
