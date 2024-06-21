package postgres

import (
	"database/sql"
	"fmt"
	"internation/model"
	"internation/pkg"
	"time"
)

type LessonRepo struct {
	Db *sql.DB
}

func NewLessonRepo(db *sql.DB) *LessonRepo {
	return &LessonRepo{db}
}

// Create
func (u *LessonRepo) CreateLesson(lesson *model.Lesson) error {
	query := `
	insert into 
	lessons(course_id, title, content)
	values($1, $2, $3)
	`
	tx, err := u.Db.Begin()
	defer tx.Commit()

	if err != nil {
		return err
	}
	_, err = tx.Exec(query, lesson.CourseId, lesson.Title, lesson.Content)

	return err
}

// Read
func (u *LessonRepo) GetLessonById(id string) (*model.Lesson, error) {
	query := `
	select 
		lesson_id, course_id, title, content, created_at, updated_at
	from
		lessons
	where 
		deleted_at is null and lesson_id = $1
	`
	Lesson := model.Lesson{}
	row := u.Db.QueryRow(query, id)
	err := row.Scan(
		&Lesson.LessonId, &Lesson.CourseId, &Lesson.Title, &Lesson.Content,
		&Lesson.CreatedAt, &Lesson.UpdatedAt)

	return &Lesson, err
}

func (u *LessonRepo) GetLessons(filter model.LessonFilter) (*[]model.Lesson, error) {
	query := `
	select 
		lesson_id, course_id, title, content, created_at, updated_at
	from
		lessons
	where 
		deleted_at is null
	`
	params := []interface{}{}
	paramCount := 1
	if filter.CourseId != nil {
		query += " and course_id = "
		pkg.AppendParamPlaceholder(&query, paramCount)
		params = append(params, *filter.CourseId)
		paramCount++
	}
	if filter.Title != nil {
		query += " and title = "
		pkg.AppendParamPlaceholder(&query, paramCount)
		params = append(params, *filter.Title)
		paramCount++
	}
	if filter.Limit != nil {
		query += " and limit = "
		pkg.AppendParamPlaceholder(&query, paramCount)
		params = append(params, *filter.Limit)
		paramCount++
	}
	if filter.Offset != nil {
		query += " and offset = "
		pkg.AppendParamPlaceholder(&query, paramCount)
		params = append(params, *filter.Offset)
		paramCount++
	}

	Lessons := []model.Lesson{}

	rows, err := u.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		Lesson := model.Lesson{}
		err = rows.Scan(
			&Lesson.LessonId, &Lesson.CourseId, &Lesson.Title, &Lesson.Content,
			&Lesson.CreatedAt, &Lesson.UpdatedAt)

		if err != nil {
			return nil, err
		}

		Lessons = append(Lessons, Lesson)
	}

	return &Lessons, rows.Err()
}

// Update
func (u *LessonRepo) UpdateLesson(lesson *model.Lesson) error {
	query := `
	update 
		lessons
	set
		course_id = $1, 
		title = $2, 
		content = $3,
		updated_at = $4
	where
		deleted_at is null and lesson_id = $5
	`
	tx, err := u.Db.Begin()
	defer tx.Commit()

	if err != nil {
		return err
	}
	result, err := tx.Exec(query, lesson.CourseId, lesson.Title, lesson.Content, time.Now(), lesson.LessonId)

	if err != nil {
		return err
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return fmt.Errorf("no rows have been affected no such lesson")
	}

	return err
}

// Delete
func (u *LessonRepo) DeleteLesson(id string) error {
	query := `
	update 
		lessons
	set
		deleted_at = $1
	where
		deleted_at is null and lesson_id = $2
	`
	tx, err := u.Db.Begin()
	defer tx.Commit()

	if err != nil {
		return err
	}
	result, err := tx.Exec(query, time.Now(), id)
	if err != nil {
		return err
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return fmt.Errorf("no rows have been affected no such user")
	}

	return nil
}
