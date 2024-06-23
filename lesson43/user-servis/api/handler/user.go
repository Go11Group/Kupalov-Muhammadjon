package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"usersevis/models"

	"github.com/gin-gonic/gin"
)

func (u *handler) CreateUser(ctx *gin.Context) {
	user := models.CreateUpdateUser{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while decoding user json",
		})
		log.Println(err)
		return
	}
	err = u.UserRepo.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while creating user",
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusCreated, "Success")
}

func (u *handler) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	if len(id) != 32 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error no id sent or invalid id in url",
		})
		return
	}

	user, err := u.UserRepo.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error":   err,
			"Message": "Error while getting user by id",
		})
		log.Println(err)
		return
	}
	ctx.JSON(200, user)
}

func (u *handler) GetUsers(ctx *gin.Context) {
	filter := models.UserFilter{}

	name := ctx.Query("name")
	if len(name) > 3 {
		filter.Name = &name
	}

	phone := ctx.Query("phone")
	if len(phone) > 3 {
		filter.Phone = &phone
	}
	ageFrom := ctx.Query("age_from")
	if len(phone) > 3 {
		age, err := strconv.Atoi(ageFrom)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error":   err,
				"Message": "Error while converting age_from to int",
			})
			log.Println(err)
			return
		}
		filter.AgeFrom = &age
	}
	ageTo := ctx.Query("age_to")
	if len(phone) > 3 {
		age, err := strconv.Atoi(ageTo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error":   err,
				"Message": "Error while converting age_to to int",
			})
			log.Println(err)
			return
		}
		filter.AgeTo = &age
	}

	users, err := u.UserRepo.GetUsers(&filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error":   err,
			"Message": "Error while getting users",
		})
		log.Println(err)
		return
	}
	ctx.JSON(200, users)
}

func (u *handler) UpdateUser(ctx *gin.Context) {
	user := models.User{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while decoding user json",
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
	user.Id = id

	err = u.UserRepo.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while updating user",
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}

func (u *handler) DeleteUser(ctx *gin.Context) {

	id := ctx.Param("id")
	if len(id) != 32 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error no id sent or invalid id in url",
		})
		return
	}

	err := u.UserRepo.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while deleting user",
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}
