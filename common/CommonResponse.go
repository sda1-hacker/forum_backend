package common

import "github.com/gin-gonic/gin"

type Res struct {
}

type Response struct {
	Code    uint        `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func (r Res) SuccessResponse(context *gin.Context, msg string, data interface{}) {
	res := Response{
		Code:    200,
		Message: msg,
		Data:    data,
	}
	context.JSON(200, res)
}

func (r Res) FailResponse(context *gin.Context, msg string, data interface{}) {
	res := Response{
		Code:    500,
		Message: msg,
		Data:    data,
	}
	context.JSON(200, res)
}

func (r Res) Unauthorized(context *gin.Context) {
	res := Response{
		Code:    401,
		Message: "请登陆之后在进行访问..",
	}
	context.JSON(200, res)
}
