package grpcbinhandler

import (
	"io"
	"math/rand"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "github.com/moul/pb/grpcbin/go-grpc"
)

type Handler struct{}

func (h *Handler) Index(ctx context.Context, in *pb.EmptyMessage) (*pb.IndexReply, error) {
	reply := pb.IndexReply{
		Description: "gRPC testing server",
		Endpoints:   []*pb.IndexReply_Endpoint{},
	}
	for _, method := range pb.GRPCBin_serviceDesc.Methods {
		reply.Endpoints = append(reply.Endpoints, &pb.IndexReply_Endpoint{
			Path: method.MethodName,
			// Description: FIXME get from comments
		})
	}
	return &reply, nil
}

func (h *Handler) HeadersUnary(ctx context.Context, in *pb.EmptyMessage) (*pb.HeadersMessage, error) {
	md, ok := metadata.FromIncomingContext(ctx)
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

func (h *Handler) NoResponseUnary(ctx context.Context, in *pb.EmptyMessage) (*pb.EmptyMessage, error) {
	return nil, nil
}

func (h *Handler) Empty(ctx context.Context, in *pb.EmptyMessage) (*pb.EmptyMessage, error) {
	return &pb.EmptyMessage{}, nil
}

func (h *Handler) DummyUnary(ctx context.Context, in *pb.DummyMessage) (*pb.DummyMessage, error) {
	return in, nil
}

func (h *Handler) RandomError(ctx context.Context, in *pb.EmptyMessage) (*pb.EmptyMessage, error) {
	c := codes.Code(uint32(rand.Intn(16)))
	return &pb.EmptyMessage{}, status.Error(c, c.String())
}

func (h *Handler) SpecificError(ctx context.Context, in *pb.SpecificErrorRequest) (*pb.EmptyMessage, error) {
	c := codes.Code(in.Code)
	msg := c.String()
	if in.Reason != "" {
		msg = in.Reason
	}
	return &pb.EmptyMessage{}, status.Error(c, msg)
}

func (h *Handler) DummyBidirectionalStreamStream(stream pb.GRPCBin_DummyBidirectionalStreamStreamServer) error {
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

func (h *Handler) DummyClientStream(stream pb.GRPCBin_DummyClientStreamServer) error {
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

func (h *Handler) DummyServerStream(in *pb.DummyMessage, stream pb.GRPCBin_DummyServerStreamServer) error {
	for i := 0; i < 10; i++ {
		if err := stream.Send(in); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
