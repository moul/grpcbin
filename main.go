//go:generate protoc -I ./grpcbin --go_out=plugins=grpc:./grpcbin ./grpcbin/grpcbin.proto

package main

import (
	"io"
	"log"
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/moul/grpcbin/grpcbin"
)

type server struct{}

func (s *server) Index(ctx context.Context, in *pb.EmptyMessage) (*pb.IndexReply, error) {
	return &pb.IndexReply{

		Description: "gRPC testing server",
		Endpoints: []*pb.IndexReply_Endpoint{
			{Path: "index", Description: "This endpoint."},
			{Path: "dummyUnary", Description: "Unary endpoint that replies a received DummyMessage."},
			{Path: "dummyClientStream", Description: "Stream endpoint that receives 10 DummyMessages and replies with the last received one."},
			{Path: "dummyServerStream", Description: "Stream endpoint that sends back 10 times the received DummyMessage."},
			{Path: "dummyBidirectionalStream", Description: "Stream endpoint that sends back a received DummyMessage indefinitely (chat mode)."},
		},
	}, nil
}

func (s *server) DummyUnary(ctx context.Context, in *pb.DummyMessage) (*pb.DummyMessage, error) {
	return in, nil
}

func (s *server) DummyBidirectionalStreamStream(stream pb.GRPCBin_DummyBidirectionalStreamStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		stream.Send(req)
	}
	return nil
}

func (s *server) DummyClientStream(stream pb.GRPCBin_DummyClientStreamServer) error {
	var req *pb.DummyMessage
	var err error
	for i := 0; i < 10; i++ {
		req, err = stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	stream.SendAndClose(req)
	return err
}

func (s *server) DummyServerStream(in *pb.DummyMessage, stream pb.GRPCBin_DummyServerStreamServer) error {
	for i := 0; i < 10; i++ {
		if err := stream.Send(in); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
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
