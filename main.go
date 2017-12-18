package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	pb "github.com/moul/grpcbin/grpcbin"
	"github.com/moul/grpcbin/handler"
)

var (
	insecureAddr = flag.String("insecure-addr", ":9000", "The ip:port combination to listen on for insecure connections")
	secureAddr   = flag.String("metrics-addr", ":9001", "The ip:port combination to listen on for secure connections")
	keyFile      = flag.String("tls-key", "cert/server.key", "TLS private key file")
	certFile     = flag.String("tls-cert", "cert/server.crt", "TLS cert file")
)

func main() {
	// parse flags
	flag.Parse()

	// insecure listener
	go func() {
		listener, err := net.Listen("tcp", *insecureAddr)
		if err != nil {
			log.Fatalf("failted to listen: %v", err)
		}

		// create gRPC server
		s := grpc.NewServer()
		pb.RegisterGRPCBinServer(s, &handler.Handlers{})
		// register reflection service on gRPC server
		reflection.Register(s)

		log.Printf("listening on %s (insecure)\n", *insecureAddr)
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// secure listener
	go func() {
		listener, err := net.Listen("tcp", *secureAddr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("failed to load TLS keys: %v", err)
		}

		// create gRPC server
		s := grpc.NewServer(grpc.Creds(creds))
		pb.RegisterGRPCBinServer(s, &handler.Handlers{})
		// register reflection service on gRPC server
		reflection.Register(s)

		log.Printf("listening on %s (secure)\n", *secureAddr)
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
	}()

	// handle Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	log.Fatalf("%s", <-c)
}
