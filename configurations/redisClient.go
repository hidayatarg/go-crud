package configurations

import "github.com/redis/go-redis/v9"

func newRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // Redis password (if any)
		DB:       0,                // Redis database index
	})
}
