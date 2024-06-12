package router

import (
	"net/http"
	"swagger/handler"

	"github.com/gin-gonic/gin"
)

func CreateServer(handler *handler.Handler) *http.Server{
	router := gin.Default()

	router.GET("/users/", handler.GetUsers)

	return &http.Server{
		Addr: ":8080",
		Handler: router,
	}
}