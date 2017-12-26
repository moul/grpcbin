package main

import (
	"flag"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	pb "github.com/moul/pb/grpcbin/go-grpc"

	"github.com/moul/grpcbin/handler"
)

var (
	insecureAddr       = flag.String("insecure-addr", ":9000", "The ip:port combination to listen on for insecure connections")
	secureAddr         = flag.String("metrics-addr", ":9001", "The ip:port combination to listen on for secure connections")
	keyFile            = flag.String("tls-key", "cert/server.key", "TLS private key file")
	certFile           = flag.String("tls-cert", "cert/server.crt", "TLS cert file")
	production         = flag.Bool("production", false, "Production mode")
	productionHTTPAddr = flag.String("production-http-addr", ":80", "The ip:port combination to listen on for production HTTP server")
)

var index = `<!DOCTYPE html>
<html>
  <body>
    <h1>grpcbin: gRPC Request & Response Service</h1>
    <h2>Endpoints</h2>
    <ul>
      <li><a href="http://grpcb.in:9000">grpc://grpcb.in:9000 (without TLS)</a></li>
      <li><a href="https://grpcb.in:9001">grpc://grpcb.in:9001 (with TLS)</a></li>
    </ul>
    <h2>Methods (<a href="https://github.com/moul/pb/blob/master/grpcbin/grpcbin.proto">grpcbin.proto</a>)</h2>
    <ul>
      {{- range .}}
      <li>{{.MethodName}}</li>
      {{- end}}
    </ul>
    <h2>Examples</h2>
    <a href="https://github.com/moul/grpcbin-example">https://github.com/moul/grpcbin-example</a>
    <h2>About</h2>
    <a href="https://github.com/moul/grpcbin">Developed</a> by <a href="https://twitter.com/moul">Manfred Touron</a>, inspired by <a href="https://httpbin.org/">https://httpbin.org/</a>
  </body>
</html>
`

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

		// serve
		log.Printf("listening on %s (insecure gRPC)\n", *insecureAddr)
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// secure listener
	go func() {
		mux := http.NewServeMux()
		t := template.New("")
		var err error
		t, err = t.Parse(index)
		if err != nil {
			log.Fatalf("failt to parse template: %v", err)
		}
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if err2 := t.Execute(w, pb.GRPCBin_serviceDesc.Methods); err != nil {
				http.Error(w, err2.Error(), http.StatusInternalServerError)
			}
		})

		// create gRPC server
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("failed to load TLS keys: %v", err)
		}
		s := grpc.NewServer(grpc.Creds(creds))
		pb.RegisterGRPCBinServer(s, &handler.Handlers{})
		// register reflection service on gRPC server
		reflection.Register(s)

		// listen and serve
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				s.ServeHTTP(w, r)
			} else {
				mux.ServeHTTP(w, r)
			}
		})
		log.Printf("listening on %s (secure gRPC + secure HTTP/2)\n", *secureAddr)
		if err := http.ListenAndServeTLS(*secureAddr, *certFile, *keyFile, handler); err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
	}()

	if *production {
		go func() {
			mux := http.NewServeMux()
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, "https://grpcb.in", 301)
			})
			log.Printf("listening on %s (production HTTP)\n", *productionHTTPAddr)
			if err := http.ListenAndServe(*productionHTTPAddr, mux); err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
		}()
	}

	// handle Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	log.Fatalf("%s", <-c)
}
