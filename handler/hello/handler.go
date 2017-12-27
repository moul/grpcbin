package hellohandler

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/context"

	pb "github.com/moul/pb/hello/go-grpc"
)

type Handler struct{}

func (h *Handler) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	greeting := "noname"
	if in.Greeting != nil {
		greeting = *in.Greeting
	}
	reply := fmt.Sprintf("hello %s", greeting)
	return &pb.HelloResponse{Reply: &reply}, nil
}

func (h *Handler) LotsOfGreetings(stream pb.HelloService_LotsOfGreetingsServer) error {
	greetings := []string{}
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		if in.Greeting != nil {
			greetings = append(greetings, *in.Greeting)
		} else {
			greetings = append(greetings, "noname")
		}
	}
	reply := fmt.Sprintf("hello %s", strings.Join(greetings, ", "))
	return stream.SendAndClose(&pb.HelloResponse{Reply: &reply})
}

func (h *Handler) LotsOfReplies(in *pb.HelloRequest, stream pb.HelloService_LotsOfRepliesServer) error {
	greeting := "noname"
	if in.Greeting != nil {
		greeting = *in.Greeting
	}
	reply := fmt.Sprintf("hello %s", greeting)

	for i := 0; i < 10; i++ {
		if err := stream.Send(&pb.HelloResponse{Reply: &reply}); err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) BidiHello(stream pb.HelloService_BidiHelloServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		greeting := "noname"
		if in.Greeting != nil {
			greeting = *in.Greeting
		}
		reply := fmt.Sprintf("hello %s", greeting)
		if err := stream.Send(&pb.HelloResponse{Reply: &reply}); err != nil {
			return err
		}
	}
	return nil
}
