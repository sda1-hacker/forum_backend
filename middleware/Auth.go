package middleware

import (
	"github.com/forum_backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	TOKEN_NAME   = "Authorization"
	TOKEN_PREFIX = "Bearer "
)

// Auth 鉴权, 大致写了一下, 逻辑再优化一下
func Auth(context *gin.Context) {
	token := context.GetHeader(TOKEN_NAME)

	// token不存在，或者token格式不正确，直接返回
	if "" == token || !strings.HasPrefix(token, TOKEN_PREFIX) {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "请在登陆之后访问..."})
		return
	}

	// token无法解析， 直接返回
	token = token[len(TOKEN_PREFIX):]
	parseToken, err := utils.JwtUtils{}.ParseToken(token)
	userId := parseToken.ID
	if err != nil || userId == 0 {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "请在登陆之后访问..."})
		return
	}

	// token续签逻辑,
	// 1. 将token持久化到redis中, 设置过期时间为token过期时间的两倍
	// 2. 使用token在redis中的有效时间作为token的过期时间
	// 3. 当redis中的tokrn过期之后, 将reids中token的过期是将向后继续延长为jwt过期时间的两倍
	// 4. 这样就可以保持前端的token不需要改变了.

	// token正确，没有过期, 放行, 将数据存储到context中，以便于其它handler使用
	context.Set("userId", parseToken.ID)
	context.Set("userName", parseToken.Name)
	context.Next() // 放行

	// context.Get("userId") 获取context存储的数据
}
