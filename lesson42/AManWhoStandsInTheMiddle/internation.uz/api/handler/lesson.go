package handler

import (
	"encoding/json"
	"internation/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllLessons(c *gin.Context) {

	filter := model.LessonFilter{}
	courseId, hasKey := c.GetQuery("course_id")
	if hasKey {
		filter.CourseId = &courseId
	}
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

	lessons, err := h.LessonRepo.GetLessons(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error in getting lessons by filter",
		})
		log.Println("Error in getting lessons by filter", err)
		return
	}
	c.JSON(200, lessons)
}

func (h *Handler) GetLessonById(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error no id",
		})
		log.Println("Error empty id param")
		return
	}

	lesson, err := h.LessonRepo.GetLessonById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error in getting Lesson by id",
		})
		log.Println("Error in getting Lesson by id", err)
		return
	}

	c.JSON(200, *lesson)
}

func (h *Handler) CreateLesson(c *gin.Context) {

	newLesson := model.Lesson{}
	err := json.NewDecoder(c.Request.Body).Decode(&newLesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while decoding lesson",
		})
		log.Println("Error while decoding lesson", err)
		return
	}

	err = h.LessonRepo.CreateLesson(&newLesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while creating new lesson",
		})
		log.Println("Error while creating new lesson ", err)
		return
	}
	c.JSON(200, "Success")
}

func (h *Handler) UpdateLesson(c *gin.Context) {

	Lesson := model.Lesson{}
	err := json.NewDecoder(c.Request.Body).Decode(&Lesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while decoding lesson",
		})
		log.Println("Error while decoding lesson", err)
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
	Lesson.LessonId = id

	err = h.LessonRepo.UpdateLesson(&Lesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while updating lesson",
		})
		log.Println("Error while updating lesson", err)
		return
	}
	c.JSON(200, "Success")
}

func (h *Handler) DeleteLesson(c *gin.Context) {

	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error no id",
		})
		log.Println("Error empty id param")
		return
	}

	err := h.LessonRepo.DeleteLesson(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad request",
			"Message": "Error while deleting lesson",
		}) 	
		log.Println("Error while deleting lesson", err)
		return
	}
	c.JSON(200, "Success")
}

