package main

import (
	pb2 "go.guide/add-grpc-service/pb"
	"go.guide/add-grpc-service/service"
	"go.guide/add-grpc-service/transport"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	hostPort string = "localhost:8002"
)

func main() {
	server := grpc.NewServer()
	sc, err := net.Listen("tcp", hostPort)
	if err != nil {
		log.Fatalf("unable to listen: %+v", err)
	}
	defer server.GracefulStop()

	pb2.RegisterAddServiceServer(server, transport.MakeAddGRPCServer(service.NewAddService()))
	_ = server.Serve(sc)
}
