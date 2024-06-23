package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateServer() *http.Server {

	h := NewHandler()

	r := gin.Default()

	users := r.Group("/users")

	users.GET("/all", h.HandleUsers)
	users.GET("/:id", h.HandleUsers)
	users.POST("/create", h.HandleUsers)
	users.PUT("/:id/update", h.HandleUsers)
	users.DELETE("/:id/delete", h.HandleUsers)

	cards := r.Group("/cards")
	cards.GET("/all", h.HandleBilling)
	cards.GET("/:id", h.HandleBilling)
	cards.POST("/create", h.HandleBilling)
	cards.PUT("/:id/update", h.HandleBilling)
	cards.DELETE("/:id/delete", h.HandleBilling)

	station := r.Group("/station")
	station.GET("/all", h.HandleBilling)
	station.GET("/:id", h.HandleBilling)
	station.POST("/create", h.HandleBilling)
	station.PUT("/:id/update", h.HandleBilling)
	station.DELETE("/:id/delete", h.HandleBilling)

	terminal := r.Group("/terminals")
	terminal.GET("/all", h.HandleBilling)
	terminal.GET("/:id", h.HandleBilling)
	terminal.POST("/create", h.HandleBilling)
	terminal.PUT("/:id/update", h.HandleBilling)
	terminal.DELETE("/:id/delete", h.HandleBilling)

	transaction := r.Group("/transactions")
	transaction.GET("/all", h.HandleBilling)
	transaction.GET("/:id", h.HandleBilling)
	transaction.POST("/create", h.HandleBilling)
	transaction.PUT("/:id/update", h.HandleBilling)
	transaction.DELETE("/:id/delete", h.HandleBilling)

	return &http.Server{
		Handler: r,
		Addr: ":7777",
	}
}
