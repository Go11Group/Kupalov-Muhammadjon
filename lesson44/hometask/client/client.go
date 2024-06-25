package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "Go11Group/Kupalov-Muhammadjon/lesson44/hometask/proto/translator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTranslatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := pb.Request{
		Words: []string{"olma"},
	}

	r, err := c.Translate(ctx, &request)
	if err != nil {
		log.Fatalf("could not translate: %v", err)
	}

	fmt.Println(r.GetTranslations())
}
