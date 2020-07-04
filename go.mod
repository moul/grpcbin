module moul.io/grpcbin

go 1.14

require (
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.9.3
	github.com/moul/pb v0.0.0-20180404114147-54bdd96e6a52
	github.com/rogpeppe/fastuuid v1.1.0 // indirect
	github.com/smartystreets/assertions v1.0.0 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/net v0.0.0-20190603091049-60506f45cf65
	golang.org/x/sys v0.0.0-20190606165138-5da285871e9c // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/grpc v1.22.0
)

replace github.com/grpc-ecosystem/grpc-gateway => github.com/moul/grpc-gateway v1.9.1-0.20190603230725-390f150e109c
