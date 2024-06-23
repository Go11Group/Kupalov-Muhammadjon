package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateServer(db *sql.DB) *http.Server{
	h := newHandler(db)

	r := gin.Default()
	users := r.Group("/users")

	users.GET("/all", h.GetUsers)
	users.GET("/:id", h.GetUserById)
	users.POST("/create", h.CreateUser)
	users.PUT("/:id/update", h.UpdateUser)
	users.DELETE("/:id/delete", h.DeleteUser)

	return &http.Server{
		Handler: r,
		Addr: ":8888",
	}
}