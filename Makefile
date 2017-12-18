install: grpcbin/grpcbin.pb.go
	go install -v .

grpcbin/grpcbin.pb.go: grpcbin/grpcbin.proto
	go generate
