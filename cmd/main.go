package main

import (
	"context"
	"fmt"

	"github.com/wrandowR/gcp-pubsub-with-redis/cmd/config"
	"github.com/wrandowR/gcp-pubsub-with-redis/internal/clients"
	"github.com/wrandowR/gcp-pubsub-with-redis/internal/entity"
	"github.com/wrandowR/gcp-pubsub-with-redis/internal/service"
)

//La idea es iniciar un server que siempre este escuchando un pub/sub de gcp
//al recibir un mensaje lo procesa y lo guarda en redis

func main() {

	ctx := context.Background()
	if err := config.ReadConfig(); err != nil {
		panic(err)
	}

	if err := clients.NewRedisClient(ctx); err != nil {
		panic(err)
	}

	GCPClient, err := clients.NewGCPClient(ctx)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	service := service.NewService(clients.RedisClient)

	ch := make(chan entity.Message)

	go func() {
		err = GCPClient.PullMsgs(ctx, ch)
		if err != nil {
			ctx.Done()
		}
	}()

	for {

		select {
		case <-ctx.Done():
			return
		case msg := <-ch:

			if err := service.ProcessMessage(ctx, &msg); err != nil {
				fmt.Println(err)
			}

		}
	}

}
