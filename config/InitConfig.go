package config

import (
	"github.com/forum_backend/logger"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Error(map[string]interface{}{"配置文件读取错误: ": err.Error()})
	}
}
