package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateServer(handler *Handler) *http.Server {
	// creating the main route and goupig them to make it more readable
	router := gin.Default()
	mainRoute := router.Group("/internation.uz")

	// User Group
	usersRoute := mainRoute.Group("/users")

	usersRoute.GET("/all", handler.Get)
	usersRoute.GET("/:id", handler.Get)
	//2ec1108a-3536-47ba-a78e-c01daf5f0a1f
	usersRoute.GET("/:id/courses", handler.Get)
	usersRoute.POST("/create", handler.Post)
	usersRoute.PUT("/:id/update", handler.Put)
	usersRoute.DELETE("/:id/delete", handler.Delete)

	// Course Group
	coursesRoute := mainRoute.Group("/courses")

	coursesRoute.GET("/all", handler.Handle)
	coursesRoute.GET("/:id", handler.Handle)
	coursesRoute.GET("/popular", handler.Handle)
	coursesRoute.GET("/:id/students", handler.Handle)
	coursesRoute.POST("/create", handler.Handle)
	coursesRoute.PUT("/:id/update", handler.Handle)
	coursesRoute.DELETE("/:id/delete", handler.Handle)

	// Enrollment Group
	enrollmentsRoute := mainRoute.Group("/enrollments")

	enrollmentsRoute.GET("/all", handler.Handle)
	enrollmentsRoute.GET("/:id", handler.Handle)
	enrollmentsRoute.POST("/create", handler.Handle)
	enrollmentsRoute.PUT("/:id/update", handler.Handle)
	enrollmentsRoute.DELETE("/:id/delete", handler.Handle)

	// Lesson Group
	lessonsRoute := mainRoute.Group("/lessons")

	lessonsRoute.GET("/all", handler.Handle)
	lessonsRoute.GET("/:id", handler.Handle)
	lessonsRoute.POST("/create", handler.Handle)
	lessonsRoute.PUT("/:id/update", handler.Handle)
	lessonsRoute.DELETE("/:id/delete", handler.Handle)

	return &http.Server{
		Addr:    ":7777",
		Handler: router,
	}
}
