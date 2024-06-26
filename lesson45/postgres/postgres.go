package postgres

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	dbname   = "library"
	user     = "postgres"
	password = "root"
)

func ConnectDB() (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port = %d dbname = %s user = %s password = %s sslmode=disable",
	host, port, dbname, user, password)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	
	return db, err
}
