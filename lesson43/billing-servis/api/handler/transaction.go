package handler

import (
	"billing_servis/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (u *handler) CreateTransaction(ctx *gin.Context) {
	Transaction := models.CreateTransaction{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&Transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while decoding Transaction json",
		})
		log.Println(err)
		return
	}
	err = u.TransactionRepo.CreateTransaction(&Transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while creating Transaction",
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusCreated, "Success")
}

func (u *handler) GetTransactionById(ctx *gin.Context) {
	id := ctx.Param("id")
	if len(id) != 32 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error no id sent or invalid id in url",
		})
		return
	}

	Transaction, err := u.TransactionRepo.GetTransactionById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error":   err,
			"Message": "Error while getting Transaction by id",
		})
		log.Println(err)
		return
	}
	ctx.JSON(200, Transaction)
}

func (u *handler) GetTransactions(ctx *gin.Context) {
	filter := models.TransactionFilter{}

	ci := ctx.Query("card_id")
	if len(ci) > 3 {
		filter.CardId = &ci
	}

	a := ctx.Query("amount")
	if len(a) > 3 {
		amount, err := strconv.Atoi(a)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error":   err,
				"Message": "Error while converting number to int",
			})
			log.Println(err)
			return
		}
		filter.Amount = &amount
	}

	ti := ctx.Query("terminal_id")
	if len(ti) > 3 {
		filter.TerminalId = &ti
	}

	Transactions, err := u.TransactionRepo.GetTransactions(&filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error":   err,
			"Message": "Error while getting Transactions",
		})
		log.Println(err)
		return
	}
	ctx.JSON(200, Transactions)
}

func (u *handler) UpdateTransaction(ctx *gin.Context) {
	Transaction := models.Transaction{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&Transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while decoding Transaction json",
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
	Transaction.Id = id

	err = u.TransactionRepo.UpdateTransaction(&Transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while updating Transaction",
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}

func (u *handler) DeleteTransaction(ctx *gin.Context) {

	id := ctx.Param("id")
	if len(id) != 32 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error no id sent or invalid id in url",
		})
		return
	}

	err := u.TransactionRepo.DeleteTransaction(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while deleting Transaction",
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}
