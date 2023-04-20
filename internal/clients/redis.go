package clients

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/wrandowR/gcp-pubsub-with-redis/internal/entity"
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

// StoreMessages stores the message in redis, #por mejorar
func StoreMessages(ctx context.Context, message *entity.Message) error {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	key := message.ID + "-key"

	result, err := RedisClient.Set(ctx, key, message.Message, 0).Result()
	if err != nil {
		return err
	}

	if result != "OK" {
		return err
	}

	return nil
}
