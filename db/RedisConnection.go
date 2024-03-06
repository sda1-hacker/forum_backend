package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func InitRedisConnection() {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	// auth := viper.GetString("redis.auth")
	db := viper.GetInt("redis.db")

	RedisClient = redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
		// Password: auth,
		DB: db,
	})
}
