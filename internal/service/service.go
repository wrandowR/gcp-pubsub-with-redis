package service

import (
	"github.com/redis/go-redis/v9"
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

func (s *Service) ProcessMessage() error {

	return nil
}
