package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type config struct {
	RedisHost string `envconfig:"REDIS_HOST" default:"localhost"`
	RedisPort string `envconfig:"REDIS_PORT" default:"6379"`
	RedisPass string `envconfig:"REDIS_PASS" default:""`
}

var c config

// ReadConfig read config
func ReadConfig() error {
	ctx := context.Background()
	err := envconfig.Process(ctx, &c)
	return err
}

// GetRedisHost returns the redis host
func GetRedisHost() string {
	return c.RedisHost
}

// GetRedisPort returns the redis port
func GetRedisPort() string {
	return c.RedisPort
}

// GetRedisPass returns the redis password
func GetRedisPass() string {
	return c.RedisPass
}
