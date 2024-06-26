package main

import (
	"context"
	"fmt"
	"net"
	"strings"
	pb "surname/genproto"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedServerServer
}

func (s *Server) GetSurname(ctx context.Context, req *pb.Request) (*pb.Surname, error){
	for _, name := range names{
		parts := strings.Split(name, " ")
		if parts[0] == req.Name{
			return &pb.Surname{Surname: parts[1]}, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func main(){
	listener, err := net.Listen("tcp", ":50055")
	if err != nil {
		panic(err)
	}
	gr := grpc.NewServer()
	pb.RegisterServerServer(gr, &Server{})

	err = gr.Serve(listener)
	if err != nil {
		panic(err)
	}
}

var names = []string{
	"Abbos Qambarov",
	"Azizbek Qobulov",
	"Bekzod Qo'chqarov",
	"Diyorbek Nematov Dadajon o'g'li",
	"Faxriddin Raximberdiyev Farxodjon o'g'li",
	"Fazliddin Xayrullayev",
	"Hamidjon Nuriddinov",
	"Hamidulloh Hamidullayev",
	"Ibrohim Umarov Fazliddin o'g'li",
	"Jamshidbek Hatamov Erkin o'g'li",
	"Javohir Abdusamatov",
	"Muhammadaziz Yoqubov",
	"Muhammadjon Ko'palov",
	"Nurmuhammad",
	"Ozodjon A'zamjonov",
	"Sanjarbek Abduraxmonov",
	"Yusupov Bobur",
	"Firdavs",
	"Ozodbek",
	"Abdulaziz Xoliqulov",
	"Intizor opa",
   }
   