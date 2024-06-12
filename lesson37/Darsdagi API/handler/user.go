package handler

import (
	"net/http"
	"strconv"
	"swagger/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsers(c *gin.Context) {
	filter := model.UserFilter{}

	FirstName, hasKey := c.GetQuery("first_name")
	if hasKey {
		filter.FirstName = &FirstName
	}
	LastName, hasKey := c.GetQuery("last_name")
	if hasKey {
		filter.LastName = &LastName
	}
	age, hasKey := c.GetQuery("age")
	if hasKey {
		age, err := strconv.Atoi(age)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"Message" : "Error converting age",
			})
		}
		filter.Age = &age
	}
	gender, hasKey := c.GetQuery("gender")
	if hasKey {
		filter.Gender = &gender
	}
	nation, hasKey := c.GetQuery("nation")
	if hasKey {
		filter.Nation = &nation
	}
	feild, hasKey := c.GetQuery("feild")
	if hasKey {
		filter.Feild = &feild
	}
	ParentName, hasKey := c.GetQuery("parent_name")
	if hasKey {
		filter.Nation = &ParentName
	}
	city, hasKey := c.GetQuery("city")
	if hasKey {
		filter.City = &city
	}
	limit, hasKey := c.GetQuery("parent_name")
	if hasKey {
		limit, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"Message" : "Error converting limit",
			})
		}
		filter.Limit = &limit
	}
	offset, hasKey := c.GetQuery("offset")
	if hasKey {
		offset, err := strconv.Atoi(offset)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"Message" : "Error converting offset",
			})
		}
		filter.Offset = &offset
	}

	users, err := h.UserRepo.GetUsers(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"Message" : "Error getting users by filter",
		})
	}
	
	c.JSON(200, *users)
}
