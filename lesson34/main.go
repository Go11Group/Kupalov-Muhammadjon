package main

import (
	"log"
	"module/handler"
	"module/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	u := postgres.NewUserRepo(db)
	p := postgres.NewProductRepo(db)
	up := postgres.NewUserProductRepo(db)
	
	h := handler.NewHandler(u, p, up)
	server := handler.CreateServer(h)

	err = server.ListenAndServe()
	if err != nil{
		panic(err)
	}

}
