package handler

import (
	"encoding/json"
	"internation/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllCourses(c *gin.Context) {
	filter := model.CourseFilter{}
	title, hasKey := c.GetQuery("title")
	if hasKey {
		filter.Title = &title
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

	courses, err := h.CourseRepo.GetCourses(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error in getting courses by filter",
		})
		log.Println("Error in getting courses by filter", err)
		return
	}
	c.JSON(200, *courses)
}

func (h *Handler) GetCourseById(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error no id",
		})
		log.Println("Error empty id param")
		return
	}

	course, err := h.CourseRepo.GetCourseById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error in getting course by id",
		})
		log.Println("Error in getting course by id", err)
		return
	}

	c.JSON(200, *course)
}

func (h *Handler) GetEnrolledUsersbyCourse(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error no id",
		})
		log.Println("Error empty id param")
		return
	}

	users, err := h.CourseRepo.GetEnrolledUsersbyCourse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error in getting users by course_id",
		})
		log.Println("Error in getting course by courseId", err)
		return
	}

	c.JSON(200, *users)
}

func (h *Handler) CreateCourse(c *gin.Context) {

	newCourse := model.Course{}
	err := json.NewDecoder(c.Request.Body).Decode(&newCourse)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while decoding course",
		})
		log.Println("Error while decoding course", err)
		return
	}

	err = h.CourseRepo.CreateCourse(&newCourse)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while creating new course",
		})
		log.Println("Error while creating new course", err)
		return
	}
	c.JSON(200, "Success")
}

func (h *Handler) UpdateCourse(c *gin.Context) {

	course := model.Course{}
	err := json.NewDecoder(c.Request.Body).Decode(&course)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while decoding course",
		})
		log.Println("Error while decoding course", err)
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

	course.CourseId = id
	err = h.CourseRepo.UpdateCourse(&course)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while updating course",
		})
		log.Println("Error while updating course", err)
		return
	}
	c.JSON(200, "Success")
}

func (h *Handler) DeleteCourse(c *gin.Context) {

	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error no id",
		})
		log.Println("Error empty id param")
		return
	}

	err := h.CourseRepo.DeleteCourse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while deleting course",
		})
		log.Println("Error while deleting course", err)
		return
	}
	c.JSON(200, "Success")
}
