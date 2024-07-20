package api

import (
	"lesson62/api/handler"
	"lesson62/api/middleware"
	"lesson62/models"

	"github.com/gin-gonic/gin"
)


func NewRouter(systemConfig *models.SystemConfig) *gin.Engine{
	router := gin.Default()

	router.Use(middleware.PermissonMiddleware(systemConfig.CasbinEnforcer))
	handler := handler.NewHandler(systemConfig)

	items := router.Group("/items")
	items.POST("/create", handler.CreateItem)
	items.GET("/:id", handler.GetItemById)
	items.GET("/all", handler.GetItems)
	items.PUT("/:id/update", handler.UpdateItem)
	items.DELETE("/:id/delete", handler.DeleteItem)

	return router
} 