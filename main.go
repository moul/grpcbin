package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/moul/grpcbin/grpcbin"
	"github.com/moul/grpcbin/handler"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failted to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGRPCBinServer(s, &handler.Handlers{})
	// register reflection service on gRPC server
	reflection.Register(s)
	log.Println("listening on :9000")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
