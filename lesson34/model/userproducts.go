package model

import (
	"database/sql"
	"time"
)

type UserProducts struct {
	Id        int
	ProductId int
	UserId    int
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
