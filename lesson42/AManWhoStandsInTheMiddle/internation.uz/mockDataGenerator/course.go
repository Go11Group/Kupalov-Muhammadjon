package mockdatagenerator

import (
	"internation/api/handler"
	"internation/model"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func InsertCourses(handler *handler.Handler) {
	courses := generateMockCourses(50)
	for _, course := range courses {
		handler.CourseRepo.CreateCourse(&course)
	}
}

var courseTitles = []string{
	"Beginner English",
	"Intermediate English",
	"Advanced English",
	"Conversational Spanish",
	"Business Spanish",
	"French for Beginners",
	"Intermediate French",
	"Advanced French",
	"German for Travelers",
	"Intermediate German",
	"Advanced German",
	"Mandarin Chinese Basics",
	"Conversational Mandarin Chinese",
	"Japanese for Beginners",
	"Intermediate Japanese",
	"Advanced Japanese",
	"Korean Language Essentials",
	"Conversational Korean",
	"Russian for Beginners",
	"Intermediate Russian",
	"Advanced Russian",
	"Italian Language Basics",
	"Portuguese for Beginners",
	"Arabic Language Introduction",
	"Hebrew for Beginners",
	"Turkish Language Basics",
	"Greek Language Essentials",
	"Swedish for Beginners",
	"Dutch Language Introduction",
	"Polish for Beginners",
	"Hindi Language Basics",
	"Bengali for Beginners",
	"Punjabi Language Essentials",
	"VietnameseInserEnrollments for Beginners",
	"Thai Language Introduction",
	"Indonesian for Beginners",
	"Malay Language Basics",
	"Tagalog for Beginners",
	"Farsi Language Introduction",
	"Urdu for Beginners",
}

func generateMockCourses(n int) []model.Course {
	var courses []model.Course

	for i := 0; i < n; i++ {
		title := courseTitles[rand.Intn(len(courseTitles))]
		description := faker.Sentence()

		course := model.Course{
			Title:       title,
			Description: description,
		}
		courses = append(courses, course)
	}
	return courses
}
