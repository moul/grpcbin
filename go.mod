module moul.io/grpcbin

go 1.14

require (
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/moul/pb v0.0.0-20220425114252-bca18df4138c
	github.com/rogpeppe/fastuuid v1.1.0 // indirect
	github.com/smartystreets/goconvey v1.8.1
	golang.org/x/crypto v0.28.0
	golang.org/x/net v0.30.0
	google.golang.org/grpc v1.45.0
)

replace github.com/grpc-ecosystem/grpc-gateway => github.com/moul/grpc-gateway v1.9.1-0.20190603230725-390f150e109c
