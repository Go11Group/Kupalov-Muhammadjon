package postgres

import (
	"database/sql"
	"fmt"
	"internation/model"
	"internation/pkg"
	"time"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

// Create
func (u *UserRepo) CreateUser(user *model.User) error {
	query := `
	insert into 
	users(name, email, birthday, password)
	values($1, $2, $3, $4)
	`
	tx, err := u.Db.Begin()
	defer tx.Commit()

	if err != nil {
		return err
	}
	_, err = tx.Exec(query, user.Name, user.Email, user.Birthday, user.Password)

	return err
}

// Read
func (u *UserRepo) GetUserById(id string) (*model.User, error) {
	query := `
	select 
		user_id, name, email, birthday, password, created_at, updated_at
	from
		users
	where 
		deleted_at is null and user_id = $1
	`
	user := model.User{}
	row := u.Db.QueryRow(query)
	err := row.Scan(
		&user.UserId, &user.Name, &user.Email, &user.Birthday, &user.Password,
		&user.CreatedAt, &user.UpdatedAt)

	return &user, err
}

func (u *UserRepo) GetUsers(filter model.UserFilter) (*[]model.User, error) {
	query := `
	select 
		user_id, name, email, birthday, password, created_at, updated_at
	from
		users
	where 
		deleted_at is null
	`
	params := []interface{}{}
	paramCount := 1
	if filter.Name != nil {
		query += " and name = "
		pkg.AppendParamPlaceholder(&query, paramCount)
		params = append(params, *filter.Name)
		paramCount++
	}
	if filter.Email != nil {
		query += " and email = "
		pkg.AppendParamPlaceholder(&query, paramCount)
		params = append(params, *filter.Email)
		paramCount++
	}
	if filter.Birthday != nil {
		query += " and birthday = "
		pkg.AppendParamPlaceholder(&query, paramCount)
		params = append(params, *filter.Birthday)
		paramCount++
	}
	if filter.AgeFrom != nil {
		query += " and extract(year from age(current_timestamp, birthday)) >= "
		pkg.AppendParamPlaceholder(&query, paramCount)
		params = append(params, *filter.AgeFrom)
		paramCount++
	}
	if filter.AgeTo != nil {
		query += " and extract(year from age(current_timestamp, birthday)) <= "
		pkg.AppendParamPlaceholder(&query, paramCount)
		params = append(params, *filter.AgeTo)
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

	users := []model.User{}

	rows, err := u.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}
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

func (u *UserRepo) GetEnrolledCoursebyUser(userId string) (*[]model.Course, error) {
	query := `
		select 
			c.course_id, c.title, c.description, c.created_at, c.updated_at
		from 
			courses as c
		join
			enrollments as e
		on
			c.course_id = e.course_id
		join
			users as u
		on
			u.user_id = e.user_id
		where
			e.user_id  = $1
	`
	rows, err := u.Db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	courses := []model.Course{}
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

	return &courses, rows.Err()
}

// Update
func (u *UserRepo) UpdateUser(user *model.User) error {
	query := `
	update 
		users
	set
		name = $1,
		email = $2,
		birthday = $3,
		password = $4,
		updated_at = $5
	where
		deleted_at is null and user_id = $6
	`
	tx, err := u.Db.Begin()
	defer tx.Commit()

	if err != nil {
		return err
	}
	result, err := tx.Exec(query, user.Name, user.Email, user.Birthday, user.Password, time.Now(), user.UserId)

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
func (u *UserRepo) DeleteUser(id string) error {
	query := `
	update 
		users
	set
		deleted_at = $1
	where
		deleted_at is null and user_id = $2
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
