package handler

import (
	"encoding/json"
	"fmt"
	"internation/model"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllUsers(c *gin.Context) {

	filter := model.UserFilter{}
	name, hasKey := c.GetQuery("name")
	if hasKey {
		filter.Name = &name
	}
	email, hasKey := c.GetQuery("email")
	if hasKey {
		filter.Email = &email
	}
	birthday, hasKey := c.GetQuery("birthday")
	if hasKey {
		parsed, err := time.Parse("2006-01-02", birthday)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error":   "Bad request",
				"Message": "Error in parsing birthday to time",
			})
			log.Println("Error in parsing birthday to time", err)
			return
		}
		filter.Birthday = &parsed
	}
	agf, hasKey := c.GetQuery("age_from")
	if hasKey {
		ageFrom, err := strconv.Atoi(agf)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error":   "Bad request",
				"Message": "Error while converting age_from",
			})
			log.Println("Error while converting age_from ", err)
			return
		}
		filter.AgeFrom = &ageFrom
	}
	agt, hasKey := c.GetQuery("age_to")
	if hasKey {
		ageTo, err := strconv.Atoi(agt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error":   "Bad request",
				"Message": "Error while converting age_to",
			})
			log.Println("Error while converting age_to ", err)
			return
		}
		filter.AgeTo = &ageTo
	}
	lms, hasKey := c.GetQuery("limit")
	if hasKey {
		limit, err := strconv.Atoi(lms)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error":   "Bad request",
				"Message": "Error while converting limit",
			})
			log.Println("Error while converting limit ", err)
			return
		}
		filter.Limit = &limit
	}
	ofs, hasKey := c.GetQuery("offset")
	if hasKey {
		offset, err := strconv.Atoi(ofs)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error":   "Bad request",
				"Message": "Error while converting offset",
			})
			log.Println("Error while converting offset ", err)
			return
		}
		filter.Offset = &offset
	}

	users, err := h.UserRepo.GetUsers(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error in getting users by filter",
		})
		log.Println("Error in getting users by filter", err)
		return
	}
	c.JSON(200, users)
}

func (h *Handler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error no id",
		})
		log.Println("Error empty id param")
		return
	}

	user, err := h.UserRepo.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error in getting user by id",
		})
		log.Println("Error in getting user by id", err)
		return
	}

	c.JSON(200, *user)
}

func (h *Handler) GetEnrolledCoursebyUser(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error no id",
		})
		log.Println("Error empty id param")
		return
	}

	courses, err := h.UserRepo.GetEnrolledCoursebyUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error in getting courses by user_id",
		})
		log.Println("Error in getting courses by user_id", err)
		return
	}

	c.JSON(200, *courses)
}

func (h *Handler) CreateUser(c *gin.Context) {

	newUser := model.User{}
	err := json.NewDecoder(c.Request.Body).Decode(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while decoding user",
		})
		log.Println("Error while decoding user", err)
		return
	}

	err = h.UserRepo.CreateUser(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while creating new user",
		})
		log.Println("Error while creating new user", err)
		return
	}
	c.JSON(200, "Success")
}

func (h *Handler) UpdateUser(c *gin.Context) {

	user := model.User{}
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while decoding user",
		})
		log.Println("Error while decoding user", err)
		return
	}
	id := c.Param("id")
	fmt.Println(id)
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error no id",
		})
		log.Println("Error empty id param")
		return
	}
	user.UserId = id
	
	err = h.UserRepo.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while updating user",
		})
		log.Println("Error while updating user", err)
		return
	}
	c.JSON(200, "Success")
}

func (h *Handler) DeleteUser(c *gin.Context) {

	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error no id",
		})
		log.Println("Error empty id param")
		return
	}

	err := h.UserRepo.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while deleting user",
		})
		log.Println("Error while deleting user", err)
		return
	}
	c.JSON(200, "Success")
}