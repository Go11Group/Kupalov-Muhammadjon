package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "market"
	password = "root"
)

func main() {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("update rasta set tur = $1 where name=$2", "meva", "Nok")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("delete from rasta where name=$1", "Bodring")
	if err != nil {
		panic(err)
	}
}
