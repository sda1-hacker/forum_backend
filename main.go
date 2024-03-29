package main

import (
	"github.com/forum_backend/config"
	"github.com/forum_backend/db"
	"github.com/forum_backend/logger"
	"github.com/forum_backend/middleware"
	"github.com/forum_backend/router"
	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化配置
	config.InitConfig()
	// 初始化mysql Gorm
	db.InitMysqlConnection()
}

func main() {
	e := gin.Default()

	// 通过中间件的方式开启日志
	e.Use(gin.LoggerWithConfig(logger.LoggerToFile())) // 访问日志
	e.Use(logger.Recover)                              // panic日志
	e.Use(middleware.Cors)                             // 允许跨域

	// 注册路由
	router.CommonRouter(e)
	router.UserRouter(e)

	// 启动项目
	err := e.Run(":9999")
	if err != nil {
		return
	}
}
