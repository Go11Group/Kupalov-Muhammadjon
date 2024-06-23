package handler

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{
	Client *http.Client
}

func NewHandler() *Handler{
	return &Handler{&http.Client{}}
}

func (h *Handler) Handle(c *gin.Context) {

	url := "http://localhost:8888" + c.Request.URL.String()

	req, err := http.NewRequest(c.Request.Method, url, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "StatusBadRequest",
			"Message": err,
		})
		log.Println("Error while making rerquest ", err)
		return
	}
	resp, err := h.Client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "StatusBadRequest",
			"Message": err,
		})
		log.Println("Error while getting rerquest ", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "StatusBadRequest",
			"Message": err,
		})
		log.Println("Bad request ", err)
		return
	}
	c.JSON(200, string(body))
}
