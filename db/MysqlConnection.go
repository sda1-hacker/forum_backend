package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var MysqlConnection *gorm.DB

func InitMysqlConnection() {
	// 打印SQL语句
	sqlLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags|log.Lshortfile|log.Ldate|log.Ltime),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
			LogLevel:      logger.Info,
		})
	// 组装sqlurl
	sqlUrl := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	host := viper.Get("mysql.host")
	port := viper.GetString("mysql.port")
	userName := viper.Get("mysql.userName")
	password := viper.Get("mysql.password")
	database := viper.Get("mysql.database")
	sqlUrl = fmt.Sprintf(sqlUrl, userName, password, host, port, database)
	println(sqlUrl)
	// 获取数据库连接
	db, err := gorm.Open(mysql.Open(sqlUrl), &gorm.Config{Logger: sqlLogger})
	if err != nil {

		log.Printf("数据库连接错误, err: %v \n", err)
		panic("数据库连接错误..")
	}
	MysqlConnection = db
}
