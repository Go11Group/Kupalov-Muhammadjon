package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error){
	dbUrl := "postgres://postgres:root@localhost:5432/media?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbUrl))
	if err != nil {
		return nil, err
	}
	return db, nil
}


