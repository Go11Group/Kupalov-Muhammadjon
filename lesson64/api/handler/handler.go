package handler

import (
	r "lesson64/storage/redis"
	redis "github.com/redis/go-redis/v9"
)

type Handler struct{
	stockRepo *r.StockRepo
}

func NewHandler(db *redis.Client) *Handler{
	return &Handler{
		stockRepo: r.NewStockRepo(db),
	}
}