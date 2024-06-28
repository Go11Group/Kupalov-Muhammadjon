package pkg

import (
	pbt "Go11Group/Kupalov-Muhammadjon/lesson47/api_gateway/genproto/TransportService"
	pbw "Go11Group/Kupalov-Muhammadjon/lesson47/api_gateway/genproto/WheatherService"
	"google.golang.org/grpc"
)

func CreateTransportServicClient(conn *grpc.ClientConn) *pbt.TransportServiceClient {
	trc := pbt.NewTransportServiceClient(conn)
	return &trc
}

func CreateWheatherServicClient(conn *grpc.ClientConn) *pbw.WheatherServiceClient {
	whc := pbw.NewWheatherServiceClient(conn)
	return &whc
}
