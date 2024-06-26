package main

import (
	"context"
	"database/sql"
	pb "library/genproto"
	"library/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

type LibraryServiceServer struct {
	pb.UnimplementedLibraryServiceServer
	Db *sql.DB
}

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("Unable to connect Database: ", err)
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterLibraryServiceServer(s, &LibraryServiceServer{Db: db})

	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}

func (l *LibraryServiceServer) AddBook(ctx context.Context, book *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	query := `
	insert into 
		books(title, author, year_published)
		values($1, $2, $3)
	`
	tx, err := l.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := tx.Exec(query, book.Title, book.Author, book.YearPublished)
}

func (l *LibraryServiceServer) SearchBook(ctx context.Context, query *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {

	books := []*pb.Book{}

	rows, err := l.Db.Query(query.Query)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		book := pb.Book{}
		err = rows.Scan(book.BookId, book.Title, book.Author, book.YearPublished)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return &pb.SearchBookResponse{Books: books}, rows.Err()
}

func (l *LibraryServiceServer) BorrowBook(ctx context.Context, br *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {

	query := `
	select 
		*
	from
		books
	where
		id = $1
	`
	rows, err := l.Db.Query(br.BookId)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		book := pb.Book{}
		err = rows.Scan(book.BookId, book.Title, book.Author, book.YearPublished)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return &pb.SearchBookResponse{Books: books}, rows.Err()
}
