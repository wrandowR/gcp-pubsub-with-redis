package service

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/wrandowR/gcp-pubsub-with-redis/internal/entity"
	//import entity
)

type Service struct {
	RedisClient *redis.Client
}

func NewService(redisClient *redis.Client) *Service {
	return &Service{
		RedisClient: redisClient,
	}
}

// ProcessMessage process the message and save it to redis, any logic with the message can be added here
func (s *Service) ProcessMessage(ctx context.Context, message *entity.Message) error {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	key := message.ID + "-key"

	result, err := s.RedisClient.Set(ctx, key, message.Message, 0).Result()
	if err != nil {
		return err
	}
	if result != "OK" {
		//por validar
		return err
	}

	return nil
}
