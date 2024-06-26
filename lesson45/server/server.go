package main

import (
	"context"
	"database/sql"
	"fmt"
	pb "library/genproto"
	"library/postgres"
	"log"
	"net"

	"github.com/google/uuid"
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

	listener, err := net.Listen("tcp", ":50055")
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
		books(id, title, author, year_published)
		values($1, $2, $3, $4)
	`
	tx, err := l.Db.Begin()
	if err != nil {
		return &pb.AddBookResponse{}, err
	}
	id := uuid.NewString()
	res, err := tx.Exec(query, id, book.Title, book.Author, book.YearPublished)
	if err != nil {
		tx.Rollback()
		return &pb.AddBookResponse{}, err
	}
	affRows, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return &pb.AddBookResponse{}, err
	}
	if affRows == 0 {
		tx.Rollback()
		return &pb.AddBookResponse{}, fmt.Errorf("nothing created")
	}

	err = tx.Commit()
	if err != nil {
		return &pb.AddBookResponse{}, fmt.Errorf("cannot commit")
	}

	return &pb.AddBookResponse{BookId: id}, nil
}

func (l *LibraryServiceServer) SearchBook(ctx context.Context, query *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {

	books := []*pb.Book{}

	rows, err := l.Db.Query(query.Query)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		book := pb.Book{}
		err = rows.Scan(&book.BookId, &book.Title, &book.Author, &book.YearPublished)
		if err != nil {
			return &pb.SearchBookResponse{}, err
		}
		books = append(books, &book)
	}

	return &pb.SearchBookResponse{Books: books}, rows.Err()
}

func (l *LibraryServiceServer) BorrowBook(ctx context.Context, br *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {

	query := `
	select 
		author
	from
		books
	where
		id = $1
	`

	book := pb.Book{}
	err := l.Db.QueryRow(query, br.BookId).Scan(&book.Author)

	if err != nil {
		return &pb.BorrowBookResponse{Success: false}, err
	}

	if book.Author == "" {
		return nil, fmt.Errorf("book doesnt exists")
	}

	return &pb.BorrowBookResponse{Success: true}, nil
}
