package addsvchandler

import (
	"golang.org/x/net/context"

	pb "github.com/moul/pb/addsvc/go-grpc"

	addservice "github.com/moul/grpcbin/handler/addsvc/imported"
)

type Handler struct{}

func (s Handler) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumReply, error) {
	v, err := addservice.New().Sum(ctx, int(in.A), int(in.B))
	errStr := ""
	if err != nil {
		errStr = err.Error()
	}
	return &pb.SumReply{V: int64(v), Err: errStr}, nil
}

func (s Handler) Concat(ctx context.Context, in *pb.ConcatRequest) (*pb.ConcatReply, error) {
	v, err := addservice.New().Concat(ctx, in.A, in.B)
	errStr := ""
	if err != nil {
		errStr = err.Error()
	}
	return &pb.ConcatReply{V: v, Err: errStr}, nil
}
