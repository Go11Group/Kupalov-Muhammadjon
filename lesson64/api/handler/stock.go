package handler

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetStockPrice(ctx *gin.Context) {
	companyName := ctx.Param("name")
	fmt.Println(companyName)

	res, err := h.stockRepo.GetStockPrice(ctx, companyName)
	if err != nil {
		log.Println("Failed to get stock price")
		ctx.JSON(500, gin.H{
			"message": "Error while getting stock",
			"Error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, res)
}
