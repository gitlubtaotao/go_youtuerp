package redis

import (
	"github.com/go-redis/redis/v7"
)

//链接redis
func Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:       "localhost:6379",
		Password:   "", // no password set
		DB:         0,  // use default DB
		MaxRetries: 2,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}

type Redis struct {
	client *redis.Client
}





