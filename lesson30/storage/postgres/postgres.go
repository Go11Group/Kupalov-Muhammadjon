package postgres

import (
	"database/sql"
	_"github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	dbUrl := "postgres://postgres:root@localhost:5432/userproduct?sslmode=disable"

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}