package postgres

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	dbName = "study"
	password = "root"
)

func ConnectDB()(*sql.DB, error){
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
	host, port, user, dbName, password)	

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}