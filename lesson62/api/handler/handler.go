package handler

import (
	"lesson62/models"
	"lesson62/storage/redis"

	"go.uber.org/zap"
)

type Handler struct {
	itemRepo     *redis.ItemRepo
	deliveryRepo *redis.DeliveryRepo
	log          *zap.Logger
}

func NewHandler(systemConfig *models.SystemConfig) *Handler {
	return &Handler{
		itemRepo:     redis.NewItemRepo(systemConfig.RedisDb),
		deliveryRepo: redis.NewDeliveryRepo(systemConfig.RedisDb),
		log:          systemConfig.Logger,
	}
}
