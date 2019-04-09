package data

import (
	"../../common"
	"../../conf"
	"github.com/go-redis/redis"
	"time"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         conf.Conf.Redis.Addr,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     100,
		PoolTimeout:  30 * time.Second,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		common.Panic4Logger(nil, "redis connect error: %v", err)
	}
}
