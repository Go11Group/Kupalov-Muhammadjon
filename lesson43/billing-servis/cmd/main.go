package main

import (
	"billing_servis/api/handler"
	"billing_servis/storage/postgres"
	"log"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	server := handler.CreateServer(db)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
