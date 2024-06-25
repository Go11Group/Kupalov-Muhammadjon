package main

import (
	pb "mod/proto/translator"
	"google.golang.org/grpc/credentials/insecure"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewTranslatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Translate(ctx, &pb.Request{Words: []string{"moshina", "olma"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.GetTranslations())
}
