package mockdatagenerator

import (
	"fmt"
	"internation/api/handler"
	"internation/model"
	"math/rand"
	"time"
)

var lessonTitles = []string{
	"Introduction to",
	"Advanced",
	"Fundamentals of",
	"Mastering",
	"Principles of",
	"Essentials of",
	"Exploring",
	"Comprehensive",
	"Basics of",
	"Advanced Topics in",
}

var lessonContents = []string{
	"This lesson covers the basics of the subject.",
	"In this lesson, we delve deeper into advanced concepts.",
	"Fundamental principles are explored in this lesson.",
	"This course will help you master the topic.",
	"This lesson provides an overview of key principles.",
	"Essential topics are covered in detail in this lesson.",
	"Explore the various aspects of the subject in this lesson.",
	"This lesson offers a comprehensive understanding of the topic.",
	"Learn the basics in this introductory lesson.",
	"Advanced topics and techniques are discussed in this lesson.",
}

func InsertLessons(handler *handler.Handler) error {
	courseIds, err := getCourseIds(handler)
	if err != nil {
		return err
	}

	if len(courseIds) == 0 {
		return fmt.Errorf("no courses available to create lessons for")
	}

	lessons := make([]model.Lesson, 0, 100)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		courseId := courseIds[rand.Intn(len(courseIds))]
		title := generateLessonTitle()
		content := generateLessonContent()

		lesson := model.Lesson{
			CourseId: courseId,
			Title:    title,
			Content:  content,
		}
		lessons = append(lessons, lesson)
	}

	// Simulate inserting lessons into a repository or database
	for _, lesson := range lessons {
		if err := handler.LessonRepo.CreateLesson(&lesson); err != nil {
			return err
		}
	}

	return nil
}

func generateLessonTitle() string {
	return lessonTitles[rand.Intn(len(lessonTitles))] + " " + randomString(5)
}

func generateLessonContent() string {
	return lessonContents[rand.Intn(len(lessonContents))]
}

// Helper function to generate a random string of given length
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
