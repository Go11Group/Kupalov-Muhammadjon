package main

import (
	"internation/handler"
	"internation/router"
	"internation/storage/postgres"
	"log"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	h := handler.NewHandler(db)
	server := router.CreateServer(h)

	server.ListenAndServe()
}
