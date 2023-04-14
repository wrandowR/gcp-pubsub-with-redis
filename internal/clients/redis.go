package clients

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func NewRedisClient(ctx context.Context) error {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //config.GetRedisHost(),
		Password: "",               //config.GetRedisPassword(), // no password set
		DB:       0,                // use default DB
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return err
	}

	RedisClient = rdb
	return nil
}
