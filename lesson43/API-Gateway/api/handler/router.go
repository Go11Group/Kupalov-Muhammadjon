package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateServer() *http.Server {

	h := NewHandler()

	r := gin.Default()

	users := r.Group("/users")

	users.GET("/all", h.Handle)
	users.GET("/:id", h.Handle)
	users.POST("/create", h.Handle)
	users.PUT("/:id/update", h.Handle)
	users.DELETE("/:id/delete", h.Handle)

	cards := r.Group("/cards")
	cards.GET("/all", h.Handle)
	cards.GET("/:id", h.Handle)
	cards.POST("/create", h.Handle)
	cards.PUT("/:id/update", h.Handle)
	cards.DELETE("/:id/delete", h.Handle)

	station := r.Group("/station")
	station.GET("/all", h.Handle)
	station.GET("/:id", h.Handle)
	station.POST("/create", h.Handle)
	station.PUT("/:id/update", h.Handle)
	station.DELETE("/:id/delete", h.Handle)

	terminal := r.Group("/terminals")
	terminal.GET("/all", h.Handle)
	terminal.GET("/:id", h.Handle)
	terminal.POST("/create", h.Handle)
	terminal.PUT("/:id/update", h.Handle)
	terminal.DELETE("/:id/delete", h.Handle)

	transaction := r.Group("/terminals")
	transaction.GET("/all", h.Handle)
	transaction.GET("/:id", h.Handle)
	transaction.POST("/create", h.Handle)
	transaction.PUT("/:id/update", h.Handle)
	transaction.DELETE("/:id/delete", h.Handle)

	return &http.Server{
		Handler: r,
		Addr: ":7777",
	}
}
