package postgres

import (
	"database/sql"
	"fmt"
	"module/model"
	"time"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

// Create
func (u *UserRepo) CreateUser(user model.User) error {
	query := `
	insert into users(username, email, password)
	values($1, $2, $3)
	`
	tr, err := u.db.Begin()
	if err != nil {
		return err
	}
	defer endTransaction(tr)
	_, err = tr.Exec(query, user.Username, user.Email, user.Password)
	tr.Commit()

	return err
}

// Read
func (u *UserRepo) GetUser(f model.FilterUser) (*[]model.User, error) {
	paramCount := 1
	params := []interface{}{}
	query := `select * from users where deleted_at is null`
	if f.Id != nil{
		params = append(params, *f.Id)
		query += fmt.Sprintf(" and id = $%d", paramCount)
		paramCount++
	}
	if f.Username != nil{
		params = append(params, *f.Username)
		query += fmt.Sprintf(" and username = $%d", paramCount)
		paramCount++
	}

	tr, err := u.db.Begin()
	if err != nil {
		return nil, err
	}
	defer endTransaction(tr)

	rows, err := tr.Query(query, params...)
	if err != nil {
		return nil, err
	}

	users := []model.User{}
	for rows.Next(){
		user := model.User{}
		err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if rows.Err() != nil{
		return nil, err
	}

	return &users, nil
}

// Update
func (u *UserRepo) UpdateUser(user model.User) error{
	query := `
	update users
	set 
		username = $1,
		email = $2,
		password = $3
		updated_at = $4
	where
		id = $5
	`
	tr, err := u.db.Begin()
	if err != nil {
		return err
	}
	defer endTransaction(tr)

	_, err = tr.Exec(query, user.Username, user.Email, user.Password, time.Now(), user.Id)
	if err != nil {
		return err
	}

	return nil
}

// Delete
func (u *UserRepo) DeleteUser(id int) error{
	query := `
	update users
	set 
		deleted_at = $1
	where
		id = $2
	`
	tr, err := u.db.Begin()
	if err != nil {
		return err
	}
	defer endTransaction(tr)

	_, err = tr.Exec(query, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

func endTransaction(tr *sql.Tx){
	tr.Commit()
}