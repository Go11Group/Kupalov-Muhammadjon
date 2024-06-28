package handler

import (
	pbt "Go11Group/Kupalov-Muhammadjon/lesson47/api_gateway/genproto/TransportService"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) ReportTrafficJam(ctx *gin.Context) {
	rw := pbt.TrafficJamRequest{}
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
func (h *Handler) GetBusSchedule(ctx *gin.Context) {
	bn := ctx.Query("bus_number")
	if len(bn) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "StatusBadRequest",
			"message": "bus_number not found",
		})
		log.Println("bus_number not found")
		return
	}
	busNumber, err := strconv.Atoi(bn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "StatusBadRequest",
			"message": err.Error(),
		})
		log.Println(err)
		return
	}
	req := pbt.BusScheduleRequest{BusNumber: int32(busNumber)}
	curWeather, err := (*h.TrClient).GetBusSchedule(context.Background(), &req)
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

func (h *Handler) TrackBusLocation(ctx *gin.Context) {
	bn := ctx.Query("bus_number")
	if len(bn) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "StatusBadRequest",
			"message": "bus_number not found",
		})
		log.Println("bus_number not found")
		return
	}
	busNumber, err := strconv.Atoi(bn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "StatusBadRequest",
			"message": err.Error(),
		})
		log.Println(err)
		return
	}
	req := pbt.BusLocationRequest{BusNumber: int32(busNumber)}
	forecast, err := (*h.TrClient).TrackBusLocation(context.Background(), &req)
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
