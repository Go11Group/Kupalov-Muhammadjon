package main

import (
	"swagger/handler"
	"swagger/router"
	"swagger/storage/postgres"
)

func main() {

	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	h := handler.NewHandler(db)
	// generator.GenerateMockData(h)

	server := router.CreateServer(h)

	server.ListenAndServe()
}
