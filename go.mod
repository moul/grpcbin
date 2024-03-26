module moul.io/grpcbin

go 1.14

require (
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1
	github.com/moul/pb v0.0.0-20180404114147-54bdd96e6a52
	github.com/smartystreets/goconvey v1.7.2
	golang.org/x/crypto v0.18.0
	golang.org/x/net v0.20.0
	google.golang.org/grpc v1.61.0
)

replace github.com/grpc-ecosystem/grpc-gateway => github.com/moul/grpc-gateway v1.9.1-0.20190603230725-390f150e109c
