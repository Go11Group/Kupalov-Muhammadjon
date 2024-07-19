package handler

import (
	"encoding/json"
	"lesson62/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) CreateItem(ctx *gin.Context){
	req := models.Item{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.log.Info("Error while decoding ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while decoding ",
			"error": err.Error(),
		})
		return
	}

	res, err := h.itemRepo.CreateItem(ctx, &req)
	if err != nil {
		h.log.Info("Error while creating item ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while creating item ",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(201, res)
} 

func (h *Handler) GetItemById(ctx *gin.Context){

	id := ctx.Param("id")
	res, err := h.itemRepo.GetItemById(ctx, id)
	if err != nil {
		h.log.Info("Error while getting item by id", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while getting item by id ",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, res)
} 

func (h *Handler) GetItems(ctx *gin.Context){

	res, err := h.itemRepo.GetItems(ctx)
	if err != nil {
		h.log.Info("Error while getting items", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while getting items ",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, res)
} 

func (h *Handler) UpdateItem(ctx *gin.Context){

	id := ctx.Param("id")

	req := models.ItemUpdate{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.log.Info("Error while decoding ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while decoding ",
			"error": err.Error(),
		})
		return
	}
	req.Id = id

	res, err := h.itemRepo.UpdateItem(ctx, &req)
	if err != nil {
		h.log.Info("Error while updating item by id", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while updating item by id ",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, res)
} 

func (h *Handler) DeleteItem(ctx *gin.Context){

	id := ctx.Param("id")
	err := h.itemRepo.DeleteItem(ctx, id)
	if err != nil {
		h.log.Info("Error while delete item by id", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while delete item by id ",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, "success")
} 