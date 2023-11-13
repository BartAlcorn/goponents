package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/bartalcorn/goponents/pkg/application"
)

func init() {
	config := application.LoadConfig()

	application.App.RedisDB = redis.NewClient(&redis.Options{
		Addr: config.RedisAddress,
	})

	fmt.Println("Connected to Redis")

}
