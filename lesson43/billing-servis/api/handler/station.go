package handler

import (
	"billing_servis/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *handler) CreateStation(ctx *gin.Context) {
	Station := models.CreateStation{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&Station)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while decoding Station json",
		})
		log.Println(err)
		return
	}
	err = u.StationRepo.CreateStation(&Station)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while creating Station",
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusCreated, "Success")
}

func (u *handler) GetStationById(ctx *gin.Context) {
	id := ctx.Param("id")
	if len(id) != 32 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error no id sent or invalid id in url",
		})
		return
	}

	Station, err := u.StationRepo.GetStationById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error":   err,
			"Message": "Error while getting Station by id",
		})
		log.Println(err)
		return
	}
	ctx.JSON(200, Station)
}

func (u *handler) GetStations(ctx *gin.Context) {
	filter := models.StationFilter{}

	n := ctx.Query("name")
	if len(n) > 3 {
		filter.Name = &n
	}

	Stations, err := u.StationRepo.GetStations(&filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error":   err,
			"Message": "Error while getting Stations",
		})
		log.Println(err)
		return
	}
	ctx.JSON(200, Stations)
}

func (u *handler) UpdateStation(ctx *gin.Context) {
	Station := models.Station{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&Station)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while decoding Station json",
		})
		log.Println(err)
		return
	}
	id := ctx.Param("id")
	if len(id) != 32 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error no id sent or invalid id in url",
		})
		return
	}
	Station.Id = id

	err = u.StationRepo.UpdateStation(&Station)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while updating Station",
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}

func (u *handler) DeleteStation(ctx *gin.Context) {

	id := ctx.Param("id")
	if len(id) != 32 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error no id sent or invalid id in url",
		})
		return
	}

	err := u.StationRepo.DeleteStation(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while deleting Station",
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}
