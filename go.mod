module moul.io/grpcbin

go 1.14

require (
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/moul/pb v0.0.0-20180404114147-54bdd96e6a52
	github.com/rogpeppe/fastuuid v1.1.0 // indirect
	github.com/smartystreets/assertions v1.0.0 // indirect
	github.com/smartystreets/goconvey v1.6.4
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110
	google.golang.org/grpc v1.36.1
)

replace github.com/grpc-ecosystem/grpc-gateway => github.com/moul/grpc-gateway v1.9.1-0.20190603230725-390f150e109c
