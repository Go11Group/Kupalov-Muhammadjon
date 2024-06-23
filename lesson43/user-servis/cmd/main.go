package main

import (
	"log"
	"usersevis/api/handler"
	"usersevis/storage/postgres"
)

func main(){

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