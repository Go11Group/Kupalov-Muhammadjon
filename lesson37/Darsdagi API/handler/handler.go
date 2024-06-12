package handler

import (
	"database/sql"
	"swagger/storage/postgres"
)

type Handler struct {
	UserRepo *postgres.UserRepo
}

func NewHandler(db *sql.DB) *Handler {
	ur := postgres.NewUserRepo(db)
	return &Handler{ur}
}
