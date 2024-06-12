package router

import (
	"internation/handler"

	"github.com/gin-gonic/gin"
)

func UsersRouter(mainRoute *gin.RouterGroup, handler *handler.Handler) {
	usersRoute := mainRoute.Group("/users")

	usersRoute.GET("/all", handler.GetAllUsers)
	usersRoute.GET("/:id", handler.GetUserById)
	usersRoute.GET("/courses/:id", handler.GetEnrolledCoursebyUser)
	usersRoute.POST("/create", handler.CreateUser)
	usersRoute.PUT("/update/:id", handler.UpdateUser)
	usersRoute.DELETE("/delete/:id", handler.DeleteUser)
}
