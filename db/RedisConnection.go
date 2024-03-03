package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var redisClient *redis.Client

func InitRedisConnection() {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	// auth := viper.GetString("redis.auth")
	db := viper.GetInt("redis.db")

	redisClient = redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
		// Password: auth,
		DB: db,
	})
}
