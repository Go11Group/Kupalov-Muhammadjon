package model

import (
	"database/sql"
	"time"
)

type Time struct {
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type DateFilter struct {
	StartDate *time.Time
	EndDate   *time.Time
	Limit     *int
	Offset    *int
}
