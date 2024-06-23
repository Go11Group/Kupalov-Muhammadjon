package main

import (
	"api_gateway/api/handler"
	"log"
)

func main() {

	server := handler.CreateServer()

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
