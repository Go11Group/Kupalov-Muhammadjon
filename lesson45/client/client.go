package main

import (
	"context"
	"fmt"
	pb "library/genproto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewLibraryServiceClient(conn)

	req := pb.AddBookRequest{
		Title: "Hamsa",
		Author: "Alisher Navoiy",
		YearPublished: 1494,
	}
	resp, err := client.AddBook(context.Background(), &req)
	if err != nil {
		log.Println("Cannot create book ", err)
	}
	fmt.Println(resp)

	search := pb.SearchBookRequest{
		Query: "select * from books",
	}
	rs, err := client.SearchBook(context.Background(), &search)
	if err != nil {
		log.Println("Cannot get books ", err)
	}
	fmt.Println(rs.Books)

	br := pb.BorrowBookRequest{
		BookId: resp.BookId,
		UserId: "456789",
	}
	brs, err := client.BorrowBook(context.Background(), &br)
	if err != nil {
		log.Println("Cannot borrow book ", err)
	}
	fmt.Println(brs)

}
