package postgres

import (
	"database/sql"
	"internation/model"
	"internation/pkg"
	"time"
)

type EnrollmentRepo struct {
	Db *sql.DB
}

func NewEnrollmentRepo(db *sql.DB) *EnrollmentRepo {
	return &EnrollmentRepo{db}
}

// Create
func (u *EnrollmentRepo) CreateEnrollment(enrollment *model.Enrollment) error {
	query := `
	insert into 
	enrollments(user_id, course_id, enrollment_date)
	values($1, $2, $3)
	`
	tx, err := u.Db.Begin()
	defer tx.Commit()

	if err != nil {
		return err
	}
	_, err = tx.Exec(query, enrollment.UserId, enrollment.CourseId, enrollment.EnrollmentDate)

	return err
}

// Read
func (u *EnrollmentRepo) GetEnrollmentById(id string) (*model.Enrollment, error) {
	query := `
	select 
		enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at
	from
		enrollments
	where 
		deleted_at is null and enrollment_id = $1
	`
	enrollment := model.Enrollment{}
	row := u.Db.QueryRow(query)
	err := row.Scan(
		&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate,
		&enrollment.CreatedAt, &enrollment.UpdatedAt)

	return &enrollment, err
}

func (u *EnrollmentRepo) GetEnrollments(filter model.EnrollmentFilter) (*[]model.Enrollment, error) {
	query := `
	select 
		enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at
	from
		enrollments
	where 
		deleted_at is null
	`
	params := []interface{}{}
	paramCount := 1
	if filter.UserId != nil {
		query += " and user_id = "
		pkg.AppendParamPlaceholder(&query, paramCount)
		params = append(params, *filter.UserId)
		paramCount++
	}
	if filter.CourseId != nil {
		query += " and course_id = "
		pkg.AppendParamPlaceholder(&query, paramCount)
		params = append(params, *filter.CourseId)
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
	

	enrollments := []model.Enrollment{}

	rows, err := u.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		enrollment := model.Enrollment{}
		err = rows.Scan(
			&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate,
			&enrollment.CreatedAt, &enrollment.UpdatedAt)

		if err != nil {
			return nil, err
		}

		enrollments = append(enrollments, enrollment)
	}

	return &enrollments, rows.Err()
}

// Update
func (u *EnrollmentRepo) UpdateEnrollment(enrollment *model.Enrollment) error {
	query := `
	update 
		enrollments
	set
		user_id = $1, 
		course_id = $2, 
		enrollment_date = $3, 
		updated_at = $4
	where
		deleted_at is null and enrollment_id = $5
	`
	tx, err := u.Db.Begin()
	defer tx.Commit()

	if err != nil {
		return err
	}
	_, err = tx.Exec(query, enrollment.UserId, enrollment.CourseId, enrollment.EnrollmentDate,
		time.Now(), enrollment.EnrollmentId)

	return err
}

// Delete
func (u *EnrollmentRepo) DeleteEnrollment(id string) error {
	query := `
	update 
		enrollments
	set
		deleted_at = $1
	where
		deleted_at is null and enrollment_id = $2
	`
	tx, err := u.Db.Begin()
	defer tx.Commit()
	
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, time.Now(), id)

	return err
}
