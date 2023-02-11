package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

// 기본 선언
var redisClient *redis.Client
var ctx = context.Background()

// GetRedis 레디스 캐시 연결
func GetRedis() *redis.Client {
	if redisClient == nil || redisClient.Ping(ctx).Err() != nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr: os.Getenv("REDIS_URL"),
			//Password: os.Getenv("REDIS_PASSWORD"),
		})
	}
	return redisClient
}
