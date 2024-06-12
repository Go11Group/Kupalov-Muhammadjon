package handler

import (
	"encoding/json"
	"internation/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllEnrollments(c *gin.Context) {

	filter := model.EnrollmentFilter{}
	userId, hasKey := c.GetQuery("user_id")
	if hasKey {
		filter.UserId = &userId
	}
	courseId, hasKey := c.GetQuery("course_id")
	if hasKey {
		filter.CourseId = &courseId
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

	Enrollments, err := h.EnrollmentRepo.GetEnrollments(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error in getting enrollments by filter",
		})
		log.Println("Error in getting enrollments by filter", err)
		return
	}
	c.JSON(200, Enrollments)
}

func (h *Handler) GetEnrollmentById(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error no id",
		})
		log.Println("Error empty id param")
		return
	}

	enrollment, err := h.EnrollmentRepo.GetEnrollmentById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error in getting Enrollment by id",
		})
		log.Println("Error in getting Enrollment by id", err)
		return
	}

	c.JSON(200, *enrollment)
}
func (h *Handler) CreateEnrollment(c *gin.Context) {

	newEnrollment := model.Enrollment{}
	err := json.NewDecoder(c.Request.Body).Decode(&newEnrollment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while decoding enrollment",
		})
		log.Println("Error while decoding enrollment", err)
		return
	}

	err = h.EnrollmentRepo.CreateEnrollment(&newEnrollment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while creating new enrollment",
		})
		log.Println("Error while creating new enrollment ", err)
		return
	}
	c.JSON(200, "Success")
}

func (h *Handler) UpdateEnrollment(c *gin.Context) {

	Enrollment := model.Enrollment{}
	err := json.NewDecoder(c.Request.Body).Decode(&Enrollment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while decoding enrollment",
		})
		log.Println("Error while decoding enrollment", err)
		return
	}

	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error no id",
		})
		log.Println("Error empty id param")
		return
	}
	Enrollment.EnrollmentId = id

	err = h.EnrollmentRepo.UpdateEnrollment(&Enrollment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while updating enrollment",
		})
		log.Println("Error while updating enrollment", err)
		return
	}
	c.JSON(200, "Success")
}

func (h *Handler) DeleteEnrollment(c *gin.Context) {

	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error no id",
		})
		log.Println("Error empty id param")
		return
	}

	err := h.EnrollmentRepo.DeleteEnrollment(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while deleting enrollment",
		})
		log.Println("Error while deleting enrollment", err)
		return
	}
	c.JSON(200, "Success")
}
