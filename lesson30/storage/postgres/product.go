package postgres

import (
	"database/sql"
	"fmt"
	"module/model"
	"time"
)

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{db}
}

// Create
func (p *ProductRepo) CreateProduct(product model.Product) error {
	query := `
	insert into products(name, description, price, stock_quantity)
	values($1, $2, $3, $4)
	`
	tr, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer endTransaction(tr)
	_, err = tr.Exec(query, product.Name, product.Description, product.Price, product.StockQuantity)
	tr.Commit()

	return err
}

// Read
func (p *ProductRepo) GetProduct(f model.FilterProducts) (*[]model.Product, error) {
	paramCount := 1
	params := []interface{}{}
	query := `select * from products where `
	if f.Id != nil {
		params = append(params, *f.Id)
		query += fmt.Sprintf("id = $%d", paramCount)
		paramCount++
	}
	if f.Name != nil {
		params = append(params, *f.Name)
		query += fmt.Sprintf(" and name = $%d", paramCount)
		paramCount++
	}
	if f.Price != nil {
		params = append(params, *f.Price)
		query += fmt.Sprintf(" and price = $%d", paramCount)
		paramCount++
	}
	if f.Description != nil {
		params = append(params, *f.Description)
		query += fmt.Sprintf(" and description = $%d", paramCount)
		paramCount++
	}
	if f.StockQuantity != nil {
		params = append(params, *f.StockQuantity)
		query += fmt.Sprintf(" and stock_quantity = $%d", paramCount)
		paramCount++
	}
	if paramCount == 1 {
		query = "select * from products"
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

	products := []model.Product{}
	for rows.Next() {
		product := model.Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.StockQuantity, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return &products, nil
}

// Update
func (p *ProductRepo) UpdateProduct(product model.Product) error {
	query := `
	update products
	set 
		name = $1,
		description = $2,
		price = $3,
		stock_quantity = $3,
	where
		id = $4
	`
	tr, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer endTransaction(tr)

	_, err = tr.Exec(query, product.Name, product.Description, product.Price, product.StockQuantity, product.Id)
	if err != nil {
		return err
	}

	return nil
}

// Delete
func (p *ProductRepo) DeleteProduct(id int) error {
	query := `
	update products
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
