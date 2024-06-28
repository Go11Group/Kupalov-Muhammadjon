package api

import (
	"Go11Group/Kupalov-Muhammadjon/lesson47/api_gateway/api/handler"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
)

func NewRouter(conn *grpc.ClientConn) *http.Server {

	handler := handler.NewHandler(conn)

	router := gin.Default()

	weather := router.Group("/weather")

	weather.GET("/current/", handler.GetCurrentWeather)
	weather.GET("/forecast/", handler.GetWeatherForecast)
	weather.POST("/report/", handler.ReportWeather)

	transport := router.Group("/transport")

	transport.GET("/busschedule/", handler.GetBusSchedule)
	transport.GET("/buslocation/", handler.TrackBusLocation)
	transport.POST("/report/", handler.ReportTrafficJam)

	return &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}
