package router

import (
	"internation/handler"

	"github.com/gin-gonic/gin"
)

func LessonsRouter(mainRoute *gin.RouterGroup, handler *handler.Handler) {
	lessonsRoute := mainRoute.Group("/lessons")
	
	lessonsRoute.GET("/all", handler.GetAllLessons)
	lessonsRoute.GET("/:id", handler.GetLessonById)
	lessonsRoute.POST("/create", handler.CreateLesson)
	lessonsRoute.PUT("/update/:id", handler.UpdateLesson)
	lessonsRoute.DELETE("/delete/:id", handler.DeleteLesson)
}
