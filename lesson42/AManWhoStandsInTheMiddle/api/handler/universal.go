package handler

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Get(c *gin.Context) {
	url := "http://localhost:8888" + c.Request.URL.String()

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "StatusBadRequest",
			"Message": err,
		})
		log.Println("Bad request ", err)
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

func (h *Handler) Post(c *gin.Context) {
	url := "http://localhost:8888" + c.Request.URL.String()
	req, err := http.NewRequest("POST", url, c.Request.Body)
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

func (h *Handler) Put(c *gin.Context) {

	url := "http://localhost:8888" + c.Request.URL.String()
	req, err := http.NewRequest("PUT", url, c.Request.Body)
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

func (h *Handler) Delete(c *gin.Context) {

	url := "http://localhost:8888" + c.Request.URL.String()
	req, err := http.NewRequest("DELETE", url, c.Request.Body)
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

// Universal hammasiga ishlaydigan code

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