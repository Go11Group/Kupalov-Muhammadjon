package mockdatagenerator

import (
	"internation/api/handler"
	"internation/model"
	"math/rand"
	"fmt"
)

func getUserIds(h *handler.Handler) ([]string, error) {
	query := `
	select
		user_id
	from
		users
	where
		deleted_at is null
	`
	rows, err := h.UserRepo.Db.Query(query)
	if err != nil {
		return []string{}, err
	}
	ids := []string{}
	for rows.Next() {
		id := ""
		err = rows.Scan(&id)
		if err != nil {
			return []string{}, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}
func getCourseIds(h *handler.Handler) ([]string, error) {
	query := `
	select
		course_id
	from
		courses
	where
		deleted_at is null
	`
	rows, err := h.UserRepo.Db.Query(query)
	if err != nil {
		return []string{}, err
	}
	ids := []string{}
	for rows.Next() {
		id := ""
		err = rows.Scan(&id)
		if err != nil {
			return []string{}, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func InsertEnrollments(handler *handler.Handler) error {
	userIds, err := getUserIds(handler)
	if err != nil {
		return err
	}
	courseIds, err := getCourseIds(handler)
	if err != nil {
		return err
	}

	enrollments := map[string]bool{}
	for i := 0; i < 100; i++ {
		var userId, courseId string
		userId = userIds[rand.Intn(len(userIds))]
		courseId = courseIds[rand.Intn(len(courseIds))]
		fmt.Println(userId)
		fmt.Println(courseId)
		if !enrollments[userId+courseId] {
			enrollments[userId+courseId] = true
			enrollment := model.Enrollment{
				UserId:   userId,
				CourseId: courseId,
			}

			if err := handler.EnrollmentRepo.CreateEnrollment(&enrollment); err != nil {
				return err
			}
		}

	}
	return nil
}
