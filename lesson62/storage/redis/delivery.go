package redis

import "github.com/redis/go-redis/v9"

type DeliveryRepo struct {
	Db *redis.Client
}

func NewDeliveryRepo(db *redis.Client) *DeliveryRepo{
	return &DeliveryRepo{Db: db}
}

