package router

import (
	"internation/handler"

	"github.com/gin-gonic/gin"
)

func CoursesRouter(mainRoute *gin.RouterGroup, handler *handler.Handler) {
	coursesRoute := mainRoute.Group("/courses")

	coursesRoute.GET("/all", handler.GetAllCourses)
	coursesRoute.GET("/:id", handler.GetCourseById)
	coursesRoute.GET("/students/:id", handler.GetEnrolledUsersbyCourse)
	coursesRoute.POST("/create", handler.CreateCourse)
	coursesRoute.PUT("/update/:id", handler.UpdateCourse)
	coursesRoute.DELETE("/delete/:id", handler.DeleteCourse)
}
