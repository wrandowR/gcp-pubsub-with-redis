package main

import (
	"context"
	"fmt"
	"time"

	"github.com/wrandowR/gcp-pubsub-with-redis/cmd/config"
	"github.com/wrandowR/gcp-pubsub-with-redis/internal/clients"
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

	for {

		select {
		case <-ctx.Done():
			return

		default:
			fmt.Println("Server running")
			//timeout
			time.Sleep(3 * time.Second)

		}

		fmt.Println("Server started")

	}
}

//Que necesito?
//1. Conectar a redis
//2. Conectar a pub/sub de gcp
//3. Procesar el mensaje
//4. Guardar el mensaje en redis
//5. Crear un server que este escuchando el pub/sub de gcp
//6. Crear un server que este escuchando el pub/sub de gcp y que este corriendo en un docker

//depronto kubernetes

//que arquitectura usar?
