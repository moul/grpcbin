//go:generate protoc -I ./grpcbin --go_out=plugins=grpc:./grpcbin ./grpcbin/grpcbin.proto

package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/moul/grpcbin/grpcbin"
)

type server struct{}

func (s *server) Index(ctx context.Context, in *pb.IndexRequest) (*pb.IndexReply, error) {
	return &pb.IndexReply{

		Description: "gRPC testing server",
		Endpoints: []*pb.IndexReply_Endpoint{
			{Path: "index", Description: "This endpoint."},
		},
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failted to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGRPCBinServer(s, &server{})
	// register reflection service on gRPC server
	reflection.Register(s)
	log.Println("listening on :8000")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
