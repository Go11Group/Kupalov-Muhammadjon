package handler

import (
	"database/sql"
	"usersevis/storage/postgres"
)

type handler struct{
	UserRepo *postgres.UserRepo
}

func newHandler(db *sql.DB) *handler{
	u := postgres.NewUserRepo(db)
	return &handler{u}
}