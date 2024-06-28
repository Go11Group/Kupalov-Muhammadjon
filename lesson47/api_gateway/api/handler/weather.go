package handler

import (
	pbw "Go11Group/Kupalov-Muhammadjon/lesson47/api_gateway/genproto/WheatherService"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) ReportWeather(ctx *gin.Context) {
	rw := pbw.ReportWheatherRequest{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&rw)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "StatusBadRequest",
			"message": "error while decoding request body",
		})
		log.Println("error while decoding request body ", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
func (h *Handler) GetCurrentWeather(ctx *gin.Context) {
	city := ctx.Query("city")
	if len(city) < 2 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "StatusBadRequest",
			"message": "City not found",
		})
		log.Println("City not found")
		return
	}
	req := pbw.CurrentWheatherRequest{City: city}
	curWeather, err := (*h.WhClient).GetCurrentWeather(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "StatusInternalServerError",
			"message": err.Error(),
		})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, curWeather)
}

func (h *Handler) GetWeatherForecast(ctx *gin.Context) {
	city := ctx.Query("city")
	if len(city) < 2 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "StatusBadRequest",
			"message": "City not found",
		})
		log.Println("City not found")
		return
	}
	d := ctx.Query("days")
	if len(d) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "StatusBadRequest",
			"message": "days not found",
		})
		log.Println("days not found")
		return
	}
	days, err := strconv.Atoi(d)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "StatusBadRequest",
			"message": err.Error(),
		})
		log.Println(err)
		return
	}
	req := pbw.ForecastWheatherRequest{City: city, Days: int32(days)}
	forecast, err := (*h.WhClient).GetWeatherForecast(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "StatusInternalServerError",
			"message": err.Error(),
		})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, forecast)
}
