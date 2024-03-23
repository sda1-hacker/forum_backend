package middleware

import (
	"github.com/gin-gonic/gin"
)

// 允许跨域中间件
func Cors(c *gin.Context) {
	// 允许跨域请求路径
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 允许跨域请求的方法
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// 允许跨域的请求头
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}
