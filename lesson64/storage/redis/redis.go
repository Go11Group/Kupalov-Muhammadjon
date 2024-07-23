package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)


func ConnectDB() (*redis.Client, error){
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	err := rdb.Ping(context.Background()).Err()

	return rdb, err
}