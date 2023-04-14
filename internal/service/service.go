package service

import (
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

func (s *Service) ProcessMessage(message entity.Message) error {

	return nil
}
