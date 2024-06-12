package postgres

import (
	"database/sql"
	"fmt"
	"internation/model"
	"internation/pkg"
	"time"
)

type CourseRepo struct {
	Db *sql.DB
}

func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{db}
}

// Create
func (c *CourseRepo) CreateCourse(course *model.Course) error {
	query := `
	insert into 
	courses(title, description)
	values($1, $2)
	`
	tx, err := c.Db.Begin()
	defer tx.Commit()

	if err != nil {
		return err
	}
	_, err = tx.Exec(query, course.Title, course.Description)

	return err
}

// Read
func (c *CourseRepo) GetCourseById(id string) (*model.Course, error) {
	query := `
	select 
		course_id, title, description, created_at, updated_at
	from
		courses
	where 
		deleted_at is null and course_id = $1
	`
	course := model.Course{}
	row := c.Db.QueryRow(query)
	err := row.Scan(
		&course.CourseId, &course.Title, &course.Description,
		&course.CreatedAt, &course.UpdatedAt)

	return &course, err
}

func (u *CourseRepo) GetCourses(filter model.CourseFilter) (*[]model.Course, error) {
	query := `
	select 
		course_id, title, description, created_at, updated_at
	from
		courses
	where 
		deleted_at is null
	`
	params := []interface{}{}
	paramCount := 1
	if filter.Title != nil {
		query += " and title = "
		pkg.AppendParamPlaceholder(&query, paramCount)
		params = append(params, *filter.Title)
		paramCount++
	}
	if filter.Limit != nil {
		query += " limit "
		pkg.AppendParamPlaceholder(&query, paramCount)
		params = append(params, *filter.Limit)
		paramCount++
	}
	if filter.Offset != nil {
		query += " offset "
		pkg.AppendParamPlaceholder(&query, paramCount)
		params = append(params, *filter.Offset)
		paramCount++
	}

	courses := []model.Course{}

	rows, err := u.Db.Query(query, params...)
	for rows.Next() {
		course := model.Course{}
		err = rows.Scan(
			&course.CourseId, &course.Title, &course.Description,
			&course.CreatedAt, &course.UpdatedAt)
		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return &courses, err
}

func (c *CourseRepo) GetEnrolledUsersbyCourse(courseId string) (*[]model.User, error) {
	query := `
		select 
			u.user_id, u.name, u.email, u.birthday, u.password, u.created_at, u.updated_at
		from 
			users as u
		join
			enrollments as e
		on
			u.user_id = e.user_id
		join
			courses as c
		on
			c.course_id = e.course_id
		where
			c.course_id  = $1
	`
	rows, err := c.Db.Query(query, courseId)
	if err != nil {
		return nil, err
	}
	users := []model.User{}
	for rows.Next() {
		user := model.User{}
		err = rows.Scan(
			&user.UserId, &user.Name, &user.Email, &user.Birthday, &user.Password,
			&user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return &users, rows.Err()
}

// Update
func (c *CourseRepo) UpdateCourse(course *model.Course) error {
	query := `
	update 
		courses
	set
		title = $1, 
		description = $2, 
		updated_at = $3
	where
		deleted_at is null and course_id = $4
	`
	tx, err := c.Db.Begin()
	defer tx.Commit()

	if err != nil {
		return err
	}
	result, err := tx.Exec(query, course.Title, course.Description, time.Now(), course.CourseId)
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
	return err
}

// Delete
func (c *CourseRepo) DeleteCourse(id string) error {
	query := `
	update 
		courses
	set
		deleted_at = $1
	where
		deleted_at is null and course_id = $2
	`
	tx, err := c.Db.Begin()
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
