package controller

import (
	"github.com/forum_backend/common"
	"github.com/forum_backend/dao"
	"github.com/forum_backend/models/dto"
	"github.com/forum_backend/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
}

func (c UserController) Login(context *gin.Context) {
	userVo := dto.LoginUserDto{}
	err := context.BindJSON(&userVo)
	userInfo, err := dao.UserDao{}.GetUserByEmail(userVo.Email)
	if err != nil {
		context.JSON(200, err.Error())
		return
	}
	if userInfo.Password == userVo.Password {
		token, _ := utils.JwtUtils{}.GenerateToken(userInfo.ID, userInfo.NickName)
		common.Res{}.SuccessResponse(context, "登陆成功..", gin.H{"userInfo": userInfo, "token": token})
		return
	} else {
		common.Res{}.FailResponse(context, "账号密码错误..", nil)
		return
	}

}

func (c UserController) AddUser(context *gin.Context) {

	// 这里是service
	// service.UserService{}.AddUser()

	context.JSON(200, gin.H{
		"message": "访问成功..",
	})
}

func (c UserController) GetUserByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		common.Res{}.FailResponse(context, "用户查询失败..", nil)
	}
	userInfo, err := dao.UserDao{}.GetUserById(uint(id))
	if err != nil {
		common.Res{}.FailResponse(context, "用户查询失败..", nil)
	} else {
		common.Res{}.SuccessResponse(context, "用户查询成功..", userInfo)
	}
}
