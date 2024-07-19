package handler

import (
	"encoding/json"
	"lesson62/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) CreateDelivery(ctx *gin.Context) {
	req := models.Delivery{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.log.Info("Error while decoding ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while decoding ",
			"error":   err.Error(),
		})
		return
	}

	res, err := h.DeliveryRepo.CreateDelivery(ctx, &req)
	if err != nil {
		h.log.Info("Error while creating Delivery ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while creating Delivery ",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(201, res)
}

func (h *Handler) GetDeliveryById(ctx *gin.Context) {

	id := ctx.Param("id")
	res, err := h.DeliveryRepo.GetDeliveryById(ctx, id)
	if err != nil {
		h.log.Info("Error while getting Delivery by id", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while getting Delivery by id ",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, res)
}

func (h *Handler) GetDeliverys(ctx *gin.Context) {

	res, err := h.DeliveryRepo.GetDeliverys(ctx)
	if err != nil {
		h.log.Info("Error while getting Deliverys", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while getting Deliverys ",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, res)
}

func (h *Handler) UpdateDelivery(ctx *gin.Context) {

	id := ctx.Param("id")

	req := models.DeliveryUpdate{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.log.Info("Error while decoding ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while decoding ",
			"error":   err.Error(),
		})
		return
	}
	req.Id = id

	res, err := h.DeliveryRepo.UpdateDelivery(ctx, &req)
	if err != nil {
		h.log.Info("Error while updating Delivery by id", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while updating Delivery by id ",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, res)
}

func (h *Handler) DeleteDelivery(ctx *gin.Context) {

	id := ctx.Param("id")
	err := h.DeliveryRepo.DeleteDelivery(ctx, id)
	if err != nil {
		h.log.Info("Error while delete Delivery by id", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while delete Delivery by id ",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, "success")
}
