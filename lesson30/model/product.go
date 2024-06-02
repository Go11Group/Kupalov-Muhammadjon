package model

import (
	"database/sql"
	"time"
)

type Product struct {
	Id            int
	Name          string
	Description   string
	Price         float32
	StockQuantity int
	CreatedAt     time.Time
	UpdatedAt     sql.NullTime
	DeletedAt     sql.NullTime
}
