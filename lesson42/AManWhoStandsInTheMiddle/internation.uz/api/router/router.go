package router

import (
	"internation/api/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateServer(handler *handler.Handler) *http.Server {
	// creating the main route and goupig them to make it more readable
	router := gin.Default()
	mainRoute := router.Group("/internation.uz")

	// User Group
	usersRoute := mainRoute.Group("/users")

	usersRoute.GET("/all/", handler.GetAllUsers)
	usersRoute.GET("/:id", handler.GetUserById)
	//2ec1108a-3536-47ba-a78e-c01daf5f0a1f
	usersRoute.GET("/:id/courses", handler.GetEnrolledCoursebyUser)
	usersRoute.POST("/create", handler.CreateUser)
	usersRoute.PUT("/:id/update", handler.UpdateUser)
	usersRoute.DELETE("/:id/delete", handler.DeleteUser)

	// Course Group
	coursesRoute := mainRoute.Group("/courses")

	coursesRoute.GET("/all/", handler.GetAllCourses)
	coursesRoute.GET("/:id", handler.GetCourseById)
	coursesRoute.GET("/popular", handler.GetPopularCourses)
	coursesRoute.GET("/:id/students", handler.GetEnrolledUsersbyCourse)
	coursesRoute.POST("/create", handler.CreateCourse)
	coursesRoute.PUT("/:id/update", handler.UpdateCourse)
	coursesRoute.DELETE("/:id/delete", handler.DeleteCourse)

	// Enrollment Group
	enrollmentsRoute := mainRoute.Group("/enrollments")

	enrollmentsRoute.GET("/all/", handler.GetAllEnrollments)
	enrollmentsRoute.GET("/:id", handler.GetEnrollmentById)
	enrollmentsRoute.POST("/create", handler.CreateEnrollment)
	enrollmentsRoute.PUT("/:id/update", handler.UpdateEnrollment)
	enrollmentsRoute.DELETE("/:id/delete", handler.DeleteEnrollment)

	// Lesson Group
	lessonsRoute := mainRoute.Group("/lessons")

	lessonsRoute.GET("/all/", handler.GetAllLessons)
	lessonsRoute.GET("/:id", handler.GetLessonById)
	lessonsRoute.POST("/create", handler.CreateLesson)
	lessonsRoute.PUT("/:id/update", handler.UpdateLesson)
	lessonsRoute.DELETE("/:id/delete", handler.DeleteLesson)

	return &http.Server{
		Addr:    ":8888",
		Handler: router,
	}
}
