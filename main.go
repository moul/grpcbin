package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	abepb "github.com/grpc-ecosystem/grpc-gateway/examples/proto/examplepb"
	addsvcpb "github.com/moul/pb/addsvc/go-grpc"
	grpcbinpb "github.com/moul/pb/grpcbin/go-grpc"
	hellopb "github.com/moul/pb/hello/go-grpc"

	abehandler "github.com/grpc-ecosystem/grpc-gateway/examples/server"
	addsvchandler "moul.io/grpcbin/handler/addsvc"
	grpcbinhandler "moul.io/grpcbin/handler/grpcbin"
	hellohandler "moul.io/grpcbin/handler/hello"
)

// Application serves webpage and grpc reqs without TLS
// Secured connections supported by proxy
var (
	host        = flag.String("host", "grpc.test.k6.io", "Host domain name")
	grpcPort    = flag.String("grpc-port", "9000", "The port for grpc server")
	webpagePort = flag.String("webpage-port", "8080", "The port for webpage HTTP server")
)

// By default proxy serves 80 and 443 for web page, 9000 and 9001 (with TLS) for grpc
var index = `<!DOCTYPE html>
<html>
  <body>
    <h1>grpcbin: gRPC Request & Response Service</h1>
    <h2>Endpoints</h2>
	<ul>
      <li>grpc://{{.Host}}:9000 (without TLS)</li>
	  <li>grpc://{{.Host}}:9001 (with TLS)</li>
	  <li><a href=http://{{.Host}}>http://{{.Host}}</a> or <a href=https://{{.Host}}>https://{{.Host}}</a> (this web page)</li>
    </ul>
    <h2>Methods</h2>
    <ul>
      <li>
        <a href="https://github.com/moul/pb/blob/master/grpcbin/grpcbin.proto">grpcbin.proto</a>
        <ul>
          {{- range .Methods}}
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
      <li>
        <a href="https://github.com/moul/pb/blob/master/a_bit_of_everything/lib/examples/examplepb/a_bit_of_everything.proto">a_bit_of_everything.proto</a>
        <ul>
          <li>Create</li>
          <li>CreateBody</li>
          <li>Lookup</li>
          <li>Update</li>
          <li>Delete</li>
          <li>GetQuery</li>
          <li>Echo</li>
          <li>DeepPathEcho</li>
          <li>NoBindings</li>
          <li>Timeout</li>
          <li>ErrorWithDetails</li>
          <li>GetMessageWithBody</li>
          <li>PostWithEmptyBody</li>
        </ul>
      </li>
    </ul>
    <h2>Examples</h2>
    <ul>
      <li><a href="https://github.com/moul/grpcbin-example">https://github.com/moul/grpcbin-example</a> (multiple languages)</li>
      <li><a href="https://github.com/lucasdicioccio/http2-client-grpc-example/">https://github.com/lucasdicioccio/http2-client-grpc-example/</a> (haskell)</li>
    </ul>
    <h2>About</h2>
	<a href="https://github.com/moul/grpcbin">Developed</a> by <a href="https://manfred.life">Manfred Touron</a>, inspired by <a href="https://httpbin.org/">https://httpbin.org/</a>
	and slightly <a href="https://github.com/loadimpact/grpcbin">tuned</a> by <a href="https://k6.io">k6.io</a>
  </body>
</html>
`

func main() {
	// parse flags
	flag.Parse()

	// grpc server
	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%s", *grpcPort))
		if err != nil {
			log.Fatalf("failted to listen: %v", err)
		}

		// create gRPC server
		s := grpc.NewServer()
		grpcbinpb.RegisterGRPCBinServer(s, &grpcbinhandler.Handler{})
		hellopb.RegisterHelloServiceServer(s, &hellohandler.Handler{})
		addsvcpb.RegisterAddServer(s, &addsvchandler.Handler{})
		abepb.RegisterABitOfEverythingServiceServer(s, abehandler.NewHandler())
		// register reflection service on gRPC server
		reflection.Register(s)

		// serve

		// if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
		// 	s.ServeHTTP(w, r)
		// }

		log.Printf("listening on %s (insecure gRPC)\n", *grpcPort)
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// webpage http server
	go func() {
		mux := http.NewServeMux()
		t := template.New("")
		var err error
		t, err = t.Parse(index)
		if err != nil {
			log.Fatalf("failt to parse template: %v", err)
		}

		webpageContent := struct {
			Host    string
			Methods []grpc.MethodDesc
		}{
			Host:    *host,
			Methods: grpcbinpb.GRPCBin_serviceDesc.Methods,
		}
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if err2 := t.Execute(w, webpageContent); err2 != nil {
				http.Error(w, err2.Error(), http.StatusInternalServerError)
			}
		})

		mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/x-icon")
			w.Header().Set("Cache-Control", "public, max-age=7776000")
			if _, err = fmt.Fprintln(w, "data:image/x-icon;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQEAYAAABPYyMiAAAABmJLR0T///////8JWPfcAAAACXBIWXMAAABIAAAASABGyWs+AAAAF0lEQVRIx2NgGAWjYBSMglEwCkbBSAcACBAAAeaR9cIAAAAASUVORK5CYII="); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})

		// listen and serve
		log.Printf("listening http on %s\n", *webpagePort)
		if err := http.ListenAndServe(fmt.Sprintf(":%s", *webpagePort), mux); err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
	}()

	// handle Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	log.Fatalf("%s", <-c)
}
