package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateServer(db *sql.DB) *http.Server {
	
	h := newHandler(db)

	r := gin.Default()

	cards := r.Group("/cards")
	cards.GET("/all", h.GetCards)
	cards.GET("/:id", h.GetCards)
	cards.POST("/create", h.CreateCard)
	cards.PUT("/:id/update", h.UpdateCard)
	cards.DELETE("/:id/delete", h.DeleteCard)

	station := r.Group("/station")
	station.GET("/all", h.GetStations)
	station.GET("/:id", h.GetStationById)
	station.POST("/create", h.CreateStation)
	station.PUT("/:id/update", h.UpdateStation)
	station.DELETE("/:id/delete", h.DeleteStation)

	terminal := r.Group("/terminals")
	terminal.GET("/all", h.GetTerminals)
	terminal.GET("/:id", h.GetTerminalById)
	terminal.POST("/create", h.CreateTerminal)
	terminal.PUT("/:id/update", h.UpdateTerminal)
	terminal.DELETE("/:id/delete", h.DeleteTerminal)

	transaction := r.Group("/transactions")
	transaction.GET("/all", h.GetTransactions)
	transaction.GET("/:id", h.GetTransactionById)
	transaction.POST("/create", h.CreateTransaction)
	transaction.PUT("/:id/update", h.UpdateTransaction)
	transaction.DELETE("/:id/delete", h.DeleteTransaction)


	return &http.Server{
		Handler: r,
		Addr: ":9999",
	}
}