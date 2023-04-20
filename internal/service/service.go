package service

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/redis/go-redis/v9"
	"github.com/wrandowR/gcp-pubsub-with-redis/internal/entity"
	//import entity
)

type Service struct {
	RedisClient *redis.Client
	GCPClient   *pubsub.Subscription
}

func NewService(redisClient *redis.Client, GCPSuscription *pubsub.Subscription) *Service {
	return &Service{
		RedisClient: redisClient,
		GCPClient:   GCPSuscription,
	}

}

func (s *Service) PullMsgs(ctx context.Context, ch chan entity.Message) error {

	err := s.GCPClient.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {

		defer msg.Ack()

		ch <- entity.Message{
			ID:      fmt.Sprint(msg.ID),
			Date:    time.Now().String(),
			Message: string(msg.Data),
		}
	})
	if err != nil {
		return fmt.Errorf("sub.Receive: %v", err)
	}

	return nil
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
