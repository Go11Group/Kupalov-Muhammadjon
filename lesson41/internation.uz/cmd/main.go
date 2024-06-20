package main

import (
	"internation/api/handler"
	"internation/api/router"
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
	// package for genrating mock data for testing purposes
	// mockdatagenerator.GenerateAll(h)

	server := router.CreateServer(h)

	server.ListenAndServe()
}
