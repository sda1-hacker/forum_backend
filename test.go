package main

//
//import (
//	"fmt"
//	"github.com/forum_backend/config"
//	"github.com/forum_backend/logger"
//	"github.com/forum_backend/utils"
//)
//
//func main() {
//	// 初始化配置
//	config.InitConfig()
//
//	token, err := utils.JwtUtils{}.GenerateToken(1, "张三")
//	if err != nil {
//		logger.Error(map[string]interface{}{})
//	}
//
//	fmt.Printf("%s \n", token)
//
//	//tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiTmFtZSI6IuW8oOS4iSIsInN1YiI6Ikp3dFRva2VuIiwiZXhwIjoxNzA5NDQ0ODQ4LCJpYXQiOjE3MDk0NDQ4MTh9.jTyygHMQCmsxHPxrlt1ycDQXPRsrIfHzWUojs3HXIRk"
//	//token, err := utils.JwtUtils{}.ParseToken(tokenStr)
//	//if err != nil {
//	//	logger.Error(map[string]interface{}{})
//	//}
//	//fmt.Printf("%v \n", token)
//}
