package postgres

import (
	"database/sql"
	"fmt"
	"time"
	"usersevis/models"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

// Create
func (u *UserRepo) CreateUser(user *models.CreateUpdateUser) error {
	query := `
	insert into
		users(name, phone, age)
		values($1, $2, $3)
	`
	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Query(query, user.Name, user.Phone, user.Age)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Read
func (u *UserRepo) GetUserById(id string) (*models.User, error) {
	user := models.User{Id: id}
	query := `
	select 
		name, phone, age
	from
		users
	where 
		id = $1 and deleted_at is null
	`

	row := u.Db.QueryRow(query, id)
	err := row.Scan(&user.Name, &user.Phone, &user.Age)
	if err != nil {
		return nil, err
	}

	return &user, row.Err()
}

func (u *UserRepo) GetUsers(filter *models.UserFilter) (*[]models.User, error) {
	query := `
	select 
		id, name, phone, age
	from
		users
	where
		deleted_at is null 
	`

	params := []interface{}{}
	paramCount := 1
	if filter.Name != nil {
		query += fmt.Sprintf("name = $%d", paramCount)
		params = append(params, *filter.Name)
		paramCount++
	}
	if filter.Phone != nil {
		query += fmt.Sprintf("phone = $%d", paramCount)
		params = append(params, *filter.Phone)
		paramCount++
	}
	if filter.AgeFrom != nil {
		query += fmt.Sprintf("age >= $%d", paramCount)
		params = append(params, *filter.AgeFrom)
		paramCount++
	}
	if filter.AgeTo != nil {
		query += fmt.Sprintf("age <= $%d", paramCount)
		params = append(params, *filter.AgeTo)
		paramCount++
	}


	rows, err := u.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Phone, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	
	return &users, rows.Err()
}

// Update
func (u *UserRepo) UpdateUser(user *models.User) error {
	query := `
	update 
		users
	set 
	`
	params := []interface{}{}
	paramCount := 1
	if len(user.Name) > 0 {
		query += fmt.Sprintf(" name = $%d", paramCount)
		params = append(params, user.Name)
		paramCount++
	}
	if len(user.Phone) > 0 {
		if paramCount > 1{
			query += ","
		}
		query += fmt.Sprintf(" phone = $%d", paramCount)
		params = append(params, user.Phone)
		paramCount++
	}
	if user.Age > 0 {
		if paramCount > 1{
			query += ","
		}
		query += fmt.Sprintf(" age = $%d", paramCount)
		params = append(params, user.Age)
		paramCount++
	}

	if paramCount > 1{
		query += ","
	}
	query += fmt.Sprintf(" updated_at = $%d", paramCount)
	params = append(params, time.Now())
	paramCount++

	query += fmt.Sprintf(" where id = $%d and deleted_at is null", paramCount)
	params = append(params, user.Id)

	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}

	result, err := tx.Exec(query, params...)
	if err != nil {
		tx.Rollback()
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if affectedRows == 0 {
		tx.Rollback()
		return fmt.Errorf("nothing updated")
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// Delete
func (u *UserRepo) DeleteUser(id string) error {
	query := `
	update 
		users
	set
		deleted_at = $1
	where
		id = $2 and deleted_at is null
	`

	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}

	result, err := tx.Exec(query,time.Now(), id)
	if err != nil {
		tx.Rollback()
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if affectedRows == 0 {
		tx.Rollback()
		return fmt.Errorf("nothing deleted")
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
