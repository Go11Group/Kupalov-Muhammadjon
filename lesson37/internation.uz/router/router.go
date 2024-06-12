package router

import (
	"internation/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateServer(handler *handler.Handler) *http.Server {
	router := gin.Default()
	mainRoute := router.Group("/internation.uz")

	UsersRouter(mainRoute, handler)
	CoursesRouter(mainRoute, handler)
	EnrollmentRouter(mainRoute, handler)
	LessonsRouter(mainRoute, handler)

	return &http.Server{
		Addr:    ":8088",
		Handler: router,
	}
}
