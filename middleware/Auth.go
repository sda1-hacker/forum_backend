package middleware

import (
	"fmt"
	"github.com/forum_backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	TOKEN_NAME   = "Authorization"
	TOKEN_PREFIX = "Bearer: "
)

// Auth 鉴权, 大致写了一下, 逻辑再优化一下
func Auth(context *gin.Context) {
	token := context.GetHeader(TOKEN_NAME)
	// token不存在，或者token格式不正确，直接返回
	if "" == token || strings.HasPrefix(token, TOKEN_PREFIX) {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}
	// token无法解析， 直接返回
	token = token[len(TOKEN_PREFIX):]
	parseToken, err := utils.ParseToken(token)
	userId := parseToken.ID
	if err != nil || userId == 0 {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}

	// token检验正确
	// 1. 与redis中的token不一致，直接返回
	// 2. redis中token过期，直接返回

	// token续签  -- redis设置的过期时间是 token的2倍，如果token过期了，redis中没有过期重新续签
	// 1. 重新生成token
	// 2. 新的token存储到redis
	// 3. 新的token返回给客户端（写到header）
	context.Header(TOKEN_NAME, fmt.Sprintf("%s%s", TOKEN_PREFIX, ""))

	// token正确，没有过期, 放行, 将数据存储到context中，以便于其它handler使用
	context.Set("userId", parseToken.ID)
	context.Set("userName", parseToken.Name)
	context.Next() // 放行

	// context.Get("userId") 获取context存储的数据
}
