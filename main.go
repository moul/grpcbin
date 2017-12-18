//go:generate protoc -I ./grpcbin --go_out=plugins=grpc:./grpcbin ./grpcbin/grpcbin.proto

package main

import (
	"io"
	"log"
	"math/rand"
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	pb "github.com/moul/grpcbin/grpcbin"
)

type server struct{}

func (s *server) Index(ctx context.Context, in *pb.EmptyMessage) (*pb.IndexReply, error) {
	return &pb.IndexReply{

		Description: "gRPC testing server",
		Endpoints: []*pb.IndexReply_Endpoint{
			{Path: "index", Description: "This endpoint."},
			{Path: "empty", Description: "Unary endpoint that takes no argument and replies an empty message."},
			{Path: "randomError", Description: "Unary endpoint that raises a random gRPC error."},
			{Path: "specificError", Description: "Unary endpoint that raises a specified (by code) gRPC error."},
			{Path: "dummyUnary", Description: "Unary endpoint that replies a received DummyMessage."},
			{Path: "dummyClientStream", Description: "Stream endpoint that receives 10 DummyMessages and replies with the last received one."},
			{Path: "dummyServerStream", Description: "Stream endpoint that sends back 10 times the received DummyMessage."},
			{Path: "dummyBidirectionalStream", Description: "Stream endpoint that sends back a received DummyMessage indefinitely (chat mode)."},
			{Path: "headers", Description: "Unary endpoint that returns headers."},
			{Path: "noResponseUnary", Description: "Unary endpoint that returns no respnose."},
		},
	}, nil
}

func (s *server) HeadersUnary(ctx context.Context, in *pb.EmptyMessage) (*pb.HeadersMessage, error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "cannot parse metadata from incoming context")
	}
	resp := pb.HeadersMessage{
		Metadata: map[string]*pb.HeadersMessage_Values{},
	}
	for key, values := range md {
		resp.Metadata[key] = &pb.HeadersMessage_Values{Values: values}
	}
	return &resp, nil
}

func (s *server) NoResponseUnary(ctx context.Context, in *pb.EmptyMessage) (*pb.EmptyMessage, error) {
	return nil, nil
}

func (s *server) Empty(ctx context.Context, in *pb.EmptyMessage) (*pb.EmptyMessage, error) {
	return &pb.EmptyMessage{}, nil
}

func (s *server) DummyUnary(ctx context.Context, in *pb.DummyMessage) (*pb.DummyMessage, error) {
	return in, nil
}

func (s *server) RandomError(ctx context.Context, in *pb.EmptyMessage) (*pb.EmptyMessage, error) {
	c := codes.Code(uint32(rand.Intn(16)))
	return &pb.EmptyMessage{}, status.Error(c, c.String())
}

func (s *server) SpecificError(ctx context.Context, in *pb.SpecificErrorRequest) (*pb.EmptyMessage, error) {
	c := codes.Code(in.Code)
	msg := c.String()
	if in.Reason != "" {
		msg = in.Reason
	}
	return &pb.EmptyMessage{}, status.Error(c, msg)
}

func (s *server) DummyBidirectionalStreamStream(stream pb.GRPCBin_DummyBidirectionalStreamStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		if err := stream.Send(req); err != nil {
			return err
		}
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
	return stream.SendAndClose(req)
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
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failted to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGRPCBinServer(s, &server{})
	// register reflection service on gRPC server
	reflection.Register(s)
	log.Println("listening on :9000")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
