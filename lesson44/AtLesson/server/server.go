package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type User struct {
	Id   int
	Name string
	Age  int
}

var userData = []User{}

type UserServer struct{}

func main() {
	userData = append(userData, User{1, "Rayim", 19})
	userServer := new(UserServer)

	rpc.Register(userServer)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	http.Serve(listener, nil)
}

func (u *UserServer) GetUserById(req *int, resp *User) error {
	fmt.Println("User getyotir")

	if *req-1 >= len(userData) || *req <= 0 {
		return fmt.Errorf("no such user")
	}

	*resp = userData[*req-1]
	fmt.Println("User getti")

	return nil
}
