package global

import "github.com/go-redis/redis/v7"

func NewRedisEngine() error {
	client := redis.NewClient(&redis.Options{
		Addr:       "localhost:6379",
		Password:   "", // no password set
		DB:         0,  // use default DB
		MaxRetries: 2,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	RedisEngine = client
	return nil
}
