package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"golang.org/x/crypto/acme/autocert"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	addsvcpb "github.com/moul/pb/addsvc/go-grpc"
	grpcbinpb "github.com/moul/pb/grpcbin/go-grpc"
	hellopb "github.com/moul/pb/hello/go-grpc"

	addsvchandler "github.com/moul/grpcbin/handler/addsvc"
	grpcbinhandler "github.com/moul/grpcbin/handler/grpcbin"
	hellohandler "github.com/moul/grpcbin/handler/hello"
)

var (
	insecureAddr       = flag.String("insecure-addr", ":9000", "The ip:port combination to listen on for insecure connections")
	secureAddr         = flag.String("metrics-addr", ":9001", "The ip:port combination to listen on for secure connections")
	keyFile            = flag.String("tls-key", "cert/server.key", "TLS private key file")
	certFile           = flag.String("tls-cert", "cert/server.crt", "TLS cert file")
	inProduction       = flag.Bool("production", false, "Production mode")
	productionHTTPAddr = flag.String("production-http-addr", ":80", "The ip:port combination to listen on for production HTTP server")
	autocertDir        = flag.String("autocert-dir", "./autocert", "Autocert (let's encrypt) caching directory")
	secureDomain       = flag.String("secure-domain", "grpcb.in", "The domain used for secure connections (ssl certificated)")
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
    <h2>Methods</h2>
    <ul>
      <li>
        <a href="https://github.com/moul/pb/blob/master/grpcbin/grpcbin.proto">grpcbin.proto</a>
        <ul>
          {{- range .}}
          <li>{{.MethodName}}</li>
          {{- end}}
        </ul>
      </li>
      <li>
        <a href="https://github.com/moul/pb/blob/master/hello/hello.proto">hello.proto</a>
        <ul>
          <li>SayHello</li>
          <li>LotsOfReplies</li>
          <li>LotsOfGreetings</li>
          <li>BidiHello</li>
        </ul>
      </li>
      <li>
        <a href="https://github.com/moul/pb/blob/master/addsvc/addsvc.proto">addsvc.proto</a>
        <ul>
          <li>Sum</li>
          <li>Concat</li>
        </ul>
      </li>
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
		grpcbinpb.RegisterGRPCBinServer(s, &grpcbinhandler.Handler{})
		hellopb.RegisterHelloServiceServer(s, &hellohandler.Handler{})
		addsvcpb.RegisterAddServer(s, &addsvchandler.Handler{})
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
		var (
			creds   credentials.TransportCredentials
			httpSrv = &http.Server{
				Addr: *secureAddr,
			}
		)

		// initialize tls configuration and grpc credentials based on production/development environment
		if *inProduction {
			hostPolicy := func(ctx context.Context, host string) error {
				allowedHost := "grpcb.in"
				if host == allowedHost {
					return nil
				}
				return fmt.Errorf("acme/autocert: only %s host is allowed", allowedHost)
			}
			m := autocert.Manager{
				Prompt:     autocert.AcceptTOS,
				HostPolicy: hostPolicy,
				Cache:      autocert.DirCache(*autocertDir),
			}
			httpSrv.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}
			creds = credentials.NewTLS(httpSrv.TLSConfig)
		} else {
			var err error
			creds, err = credentials.NewServerTLSFromFile(*certFile, *keyFile)
			if err != nil {
				log.Fatalf("failed to load TLS keys: %v", err)
			}
			cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
			if err != nil {
				log.Fatalf("failed to laod TLS keys: %v", err)
			}
			httpSrv.TLSConfig = &tls.Config{Certificates: []tls.Certificate{cert}}
		}

		// setup grpc servef
		s := grpc.NewServer(grpc.Creds(creds))
		grpcbinpb.RegisterGRPCBinServer(s, &grpcbinhandler.Handler{})
		hellopb.RegisterHelloServiceServer(s, &hellohandler.Handler{})
		addsvcpb.RegisterAddServer(s, &addsvchandler.Handler{})
		// register reflection service on gRPC server
		reflection.Register(s)

		// initilaize HTTP routing based on production/development environment
		if *inProduction {
			mux := http.NewServeMux()
			t := template.New("")
			var err error
			t, err = t.Parse(index)
			if err != nil {
				log.Fatalf("failt to parse template: %v", err)
			}
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				if err2 := t.Execute(w, grpcbinpb.GRPCBin_serviceDesc.Methods); err != nil {
					http.Error(w, err2.Error(), http.StatusInternalServerError)
				}
			})
			httpSrv.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
					s.ServeHTTP(w, r)
				} else {
					mux.ServeHTTP(w, r)
				}
			})
		} else {
			httpSrv.Handler = s
		}

		// listen and serve
		log.Printf("listening on %s (secure gRPC + secure HTTP/2)\n", *secureAddr)
		if err := httpSrv.ListenAndServeTLS("", ""); err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
	}()

	if *inProduction {
		// production HTTP server (redirect to https)
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
