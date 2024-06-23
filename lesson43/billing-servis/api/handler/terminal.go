package handler

import (
	"billing_servis/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *handler) CreateTerminal(ctx *gin.Context) {
	Terminal := models.CreateTerminal{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&Terminal)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while decoding Terminal json",
		})
		log.Println(err)
		return
	}
	err = u.TerminalRepo.CreateTerminal(&Terminal)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while creating Terminal",
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusCreated, "Success")
}

func (u *handler) GetTerminalById(ctx *gin.Context) {
	id := ctx.Param("id")
	if len(id) != 32 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error no id sent or invalid id in url",
		})
		return
	}

	Terminal, err := u.TerminalRepo.GetTerminalById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error":   err,
			"Message": "Error while getting Terminal by id",
		})
		log.Println(err)
		return
	}
	ctx.JSON(200, Terminal)
}

func (u *handler) GetTerminals(ctx *gin.Context) {
	filter := models.TerminalFilter{}

	id := ctx.Query("station_id")
	if len(id) > 3 {
		filter.StationId = &id
	}

	Terminals, err := u.TerminalRepo.GetTerminals(&filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error":   err,
			"Message": "Error while getting Terminals",
		})
		log.Println(err)
		return
	}
	ctx.JSON(200, Terminals)
}

func (u *handler) UpdateTerminal(ctx *gin.Context) {
	Terminal := models.Terminal{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&Terminal)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while decoding Terminal json",
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
	Terminal.Id = id

	err = u.TerminalRepo.UpdateTerminal(&Terminal)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while updating Terminal",
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}

func (u *handler) DeleteTerminal(ctx *gin.Context) {

	id := ctx.Param("id")
	if len(id) != 32 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error no id sent or invalid id in url",
		})
		return
	}

	err := u.TerminalRepo.DeleteTerminal(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while deleting Terminal",
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}
