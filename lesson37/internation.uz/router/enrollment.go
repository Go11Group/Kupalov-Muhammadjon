package router

import (
	"internation/handler"

	"github.com/gin-gonic/gin"
)

func EnrollmentRouter(mainRoute *gin.RouterGroup, handler *handler.Handler) {
	enrollmentsRoute := mainRoute.Group("/enrollments")
	
	enrollmentsRoute.GET("/all", handler.GetAllEnrollments)
	enrollmentsRoute.GET("/:id", handler.GetEnrollmentById)
	enrollmentsRoute.POST("/create", handler.CreateEnrollment)
	enrollmentsRoute.PUT("/update/:id", handler.UpdateEnrollment)
	enrollmentsRoute.DELETE("/delete/:id", handler.DeleteEnrollment)
}
