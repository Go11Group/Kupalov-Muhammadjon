package main

import (
	"log"
	"net/rpc"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main(){
	req := 1

	user := User{}

	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	err = client.Call("UserServer.GetUserById", req, &user)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(user)
}