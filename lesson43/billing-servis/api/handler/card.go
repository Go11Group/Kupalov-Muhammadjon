package handler

import (
	"billing_servis/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (u *handler) CreateCard(ctx *gin.Context) {
	Card := models.CreateCard{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&Card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while decoding Card json",
		})
		log.Println(err)
		return
	}
	err = u.CardRepo.CreateCard(&Card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while creating Card",
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusCreated, "Success")
}

func (u *handler) GetCardById(ctx *gin.Context) {
	id := ctx.Param("id")
	if len(id) != 32 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error no id sent or invalid id in url",
		})
		return
	}

	Card, err := u.CardRepo.GetCardById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error":   err,
			"Message": "Error while getting Card by id",
		})
		log.Println(err)
		return
	}
	ctx.JSON(200, Card)
}

func (u *handler) GetCards(ctx *gin.Context) {
	filter := models.CardFilter{}

	n := ctx.Query("number")
	if len(n) > 3 {
		num, err := strconv.Atoi(n)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error":   err,
				"Message": "Error while converting number to int",
			})
			log.Println(err)
			return
		}
		filter.Number = &num
	}

	ui := ctx.Query("user_id")
	if len(ui) > 3 {
		filter.UserId = &ui
	}

	Cards, err := u.CardRepo.GetCards(&filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error":   err,
			"Message": "Error while getting Cards",
		})
		log.Println(err)
		return
	}
	ctx.JSON(200, Cards)
}

func (u *handler) UpdateCard(ctx *gin.Context) {
	Card := models.Card{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&Card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while decoding Card json",
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
	Card.Id = id

	err = u.CardRepo.UpdateCard(&Card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while updating Card",
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}

func (u *handler) DeleteCard(ctx *gin.Context) {

	id := ctx.Param("id")
	if len(id) != 32 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error no id sent or invalid id in url",
		})
		return
	}

	err := u.CardRepo.DeleteCard(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while deleting Card",
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}
