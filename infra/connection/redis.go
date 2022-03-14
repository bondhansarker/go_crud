package connection

import (
	"fmt"

	"demo/infra/config"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func ConnectRedis() {
	redisConfig := config.Redis()

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		Password: redisConfig.Pass,
		DB:       redisConfig.Db,
	})

	if res, err := redisClient.Ping().Result(); err != nil {
		panic(err)
	} else {
		fmt.Println(res, "Redis is connected now")
	}

}

func Redis() *redis.Client {
	return redisClient
}
