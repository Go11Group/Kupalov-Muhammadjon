package mockdatagenerator

import "internation/api/handler"

// runs all mock data generator funcs
func GenerateAll(handler *handler.Handler) {
	// InsertUsers(handler)
	// InsertCourses(handler)
	InsertEnrollments(handler)
	// InsertLessons(handler)
}
