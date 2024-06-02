package postgres

import (
	"database/sql"
	"fmt"
	"module/model"
	"time"
)

type UserProductRepo struct {
	db *sql.DB
}

func NewUserProductRepo(db *sql.DB) *UserProductRepo {
	return &UserProductRepo{db}
}

// Create
func (p *UserProductRepo) CreateUserProduct(product_id, user_id int) error {
	query := `
	insert into user_products(product_id, user_id)
	values($1, $2)
	`
	tr, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer endTransaction(tr)
	_, err = tr.Exec(query, product_id, user_id)
	tr.Commit()

	return err
}

// Read
func (p *UserProductRepo) GetUserProduct(f model.FilterUserProducts) (*[]model.UserProducts, error) {
	paramCount := 1
	params := []interface{}{}
	query := `select * from user_products where `
	if f.Id != nil {
		params = append(params, *f.Id)
		query += fmt.Sprintf("id = $%d", paramCount)
		paramCount++
	}
	if f.ProductId != nil {
		params = append(params, *f.ProductId)
		query += fmt.Sprintf(" and product_id = $%d", paramCount)
		paramCount++
	}
	if f.UserId != nil {
		params = append(params, *f.UserId)
		query += fmt.Sprintf(" and user_id = $%d", paramCount)
		paramCount++
	}

	if paramCount == 1 {
		query = "select * from user_products"
	}

	tr, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	defer endTransaction(tr)

	rows, err := tr.Query(query, params...)
	if err != nil {
		return nil, err
	}

	userProducts := []model.UserProducts{}
	for rows.Next() {
		product := model.UserProducts{}
		err = rows.Scan(&product.Id, &product.ProductId, &product.UserId)
		if err != nil {
			return nil, err
		}
		userProducts = append(userProducts, product)
	}
	return &userProducts, nil
}

// Update
func (p *UserProductRepo) UpdateUserProduct(up model.UserProducts) error {
	query := `
	update user_products
	set 
		user_id = $1,
		product_id = $2,
	where
		id = $3
	`
	tr, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer endTransaction(tr)

	_, err = tr.Exec(query, up.ProductId, up.UserId, up.Id)
	if err != nil {
		return err
	}

	return nil
}

// Delete
func (p *UserProductRepo) DeleteUserProduct(id int) error {
	query := `
	update user_products
	set 
		deleted_at = $1,
	where
		id = $2
	`
	tr, err := p.db.Begin()
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
