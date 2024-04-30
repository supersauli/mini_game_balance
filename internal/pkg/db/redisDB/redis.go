package redisDB

import (
	"context"
	"github.com/go-redis/redis/v8"
	"mini_game_balance/configs"
)

func NewRedis(r *configs.Redis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password, // no password set
		DB:       r.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
