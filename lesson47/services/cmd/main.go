package main

import (
	pbt "Go11Group/Kupalov-Muhammadjon/lesson47/services/genproto/TransportService"
	pbw "Go11Group/Kupalov-Muhammadjon/lesson47/services/genproto/WheatherService"
	"Go11Group/Kupalov-Muhammadjon/lesson47/services/service"
	"Go11Group/Kupalov-Muhammadjon/lesson47/services/storage/postgres"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	server := grpc.NewServer()
	ws := service.NewWheatherService(db)
	trs := service.NewTransportService(db)

	pbw.RegisterWheatherServiceServer(server, ws)
	pbt.RegisterTransportServiceServer(server, trs)

	log.Println("Server listening")
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
