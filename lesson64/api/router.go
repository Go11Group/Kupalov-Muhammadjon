package api

import (
	"lesson64/api/handler"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)


func NewRouter(db *redis.Client) *gin.Engine{
	h := handler.NewHandler(db)

	router := gin.Default()

	st := router.Group("/stocks")

	st.GET("/:name", h.GetStockPrice)

	return router
}
