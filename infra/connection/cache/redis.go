package cache

import (
	"context"
	"fmt"

	"demo/infra/config"
	"github.com/go-redis/redis/v8"
)

type RedisCacheClient struct {
	Redis *redis.Client
}

func ConnectRedis() {
	redisConfig := config.Redis()

	client.Redis = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		Password: redisConfig.Pass,
		DB:       redisConfig.Db,
	})

	if res, err := client.Redis.Ping(context.Background()).Result(); err != nil {
		panic(err)
	} else {
		fmt.Println(res, "Redis is connected now")
	}

}
