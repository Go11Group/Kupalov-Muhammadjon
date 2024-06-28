package handler

import (
	pbt "Go11Group/Kupalov-Muhammadjon/lesson47/api_gateway/genproto/TransportService"
	pbw "Go11Group/Kupalov-Muhammadjon/lesson47/api_gateway/genproto/WheatherService"
	"Go11Group/Kupalov-Muhammadjon/lesson47/api_gateway/pkg"
	"google.golang.org/grpc"
)

type Handler struct {
	TrClient *pbt.TransportServiceClient
	WhClient *pbw.WheatherServiceClient
}

func NewHandler(conn *grpc.ClientConn) *Handler {
	whc := pkg.CreateWheatherServiceClient(conn)
	trc := pkg.CreateTransportServiceClient(conn)

	return &Handler{trc, whc}
}
