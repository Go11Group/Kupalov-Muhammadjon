package main

import (
	"context"
	"fmt"
	pb "surname/genproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	gen := pb.NewServerClient(conn)
	req := pb.Request{
		Name: "Abbos",
	}
	resp, err := gen.GetSurname(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
