package handler

import (
	"database/sql"
	"internation/storage/postgres"
)

type Handler struct {
	UserRepo       *postgres.UserRepo
	CourseRepo     *postgres.CourseRepo
	LessonRepo     *postgres.LessonRepo
	EnrollmentRepo *postgres.EnrollmentRepo
}

func NewHandler(db *sql.DB) *Handler {
	u := postgres.NewUserRepo(db)
	c := postgres.NewCourseRepo(db)
	l := postgres.NewLessonRepo(db)
	e := postgres.NewEnrollmentRepo(db)

	return &Handler{
		UserRepo:       u,
		CourseRepo:     c,
		LessonRepo:     l,
		EnrollmentRepo: e,
	}
}
