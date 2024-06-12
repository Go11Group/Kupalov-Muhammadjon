package postgres

import (
	"database/sql"
	"fmt"
	"swagger/model"
)

type UserRepo struct{
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo{
	return &UserRepo{db}
}

func (u *UserRepo) CreateUser(user model.User) error {
	query := `
		INSERT INTO users (first_name, last_name, age, gender, nation, feild, parent_name, city)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, user.FirstName, user.LastName, user.Age, user.Gender,
		user.Nation, user.Feild, user.ParentName, user.City)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit() 
	if err != nil {
		tx.Rollback() 
		return err
	}

	return nil
}


func (u *UserRepo) GetUsers(filter model.UserFilter) (*[]model.User, error){
	query := `
	select * 
	from 
		users
	where
		true`
	paramCount := 1
	params := []interface{}{}

	if filter.FirstName != nil{
		query += getFormattedString(" and first_name = ", paramCount)
		params = append(params, *filter.FirstName)
		paramCount++
	}
	if filter.LastName != nil{
		query += getFormattedString(" and last_name = ", paramCount)
		params = append(params, *filter.LastName)
		paramCount++
	}
	if filter.Age != nil{
		query += getFormattedString(" and age = ", paramCount)
		params = append(params, *filter.Age)
		paramCount++
	}
	if filter.Gender != nil{
		query += getFormattedString(" and gender = ", paramCount)
		params = append(params, *filter.Gender)
		paramCount++
	}
	if filter.Nation != nil{
		query += getFormattedString(" and nation = ", paramCount)
		params = append(params, *filter.Nation)
		paramCount++
	}
	if filter.Feild != nil{
		query += getFormattedString(" and feild = ", paramCount)
		params = append(params, *filter.Feild)
		paramCount++
	}
	if filter.ParentName != nil{
		query += getFormattedString(" and parent_name = ", paramCount)
		params = append(params, *filter.ParentName)
		paramCount++
	}
	if filter.City != nil{
		query += getFormattedString(" and city = ", paramCount)
		params = append(params, *filter.City)
		paramCount++
	}
	if filter.Limit != nil{
		query += getFormattedString(" and limit ", paramCount)
		params = append(params, *filter.Limit)
		paramCount++
	}
	if filter.Offset != nil{
		query += getFormattedString(" and offset ", paramCount)
		params = append(params, *filter.City)
		paramCount++
	}
	
	users := []model.User{}

	rows, err := u.db.Query(query, params...)
	if err != nil {
		return nil, err
	} 

	for rows.Next(){
		user := model.User{}
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Age, &user.Gender,
		&user.Nation, &user.Feild, &user.ParentName, &user.City)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return &users, rows.Err()
}

func getFormattedString(s string, paramCount int) string{
	return fmt.Sprintf(s + "$%d", paramCount)
}