package main

import (
	"Go11Group/Kupalov-Muhammadjon/lesson47/api_gateway/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.NewClient("localhost:5555", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	server := api.NewRouter(conn)

	log.Println("Starting server...")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
