module github.com/moul/grpcbin

go 1.12

require (
	github.com/grpc-ecosystem/grpc-gateway v1.9.3
	github.com/moul/pb v0.0.0-20180404114147-54bdd96e6a52
	github.com/rogpeppe/fastuuid v1.1.0 // indirect
	github.com/smartystreets/assertions v1.0.0 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a
	golang.org/x/crypto v0.0.0-20190707102051-4def268fd1a4
	golang.org/x/net v0.0.0-20190707102051-da137c7871d7
	golang.org/x/sync v0.0.0-20190603230018-112230192c58 // indirect
	golang.org/x/sys v0.0.0-20190707102051-04f50cda93cb // indirect
	google.golang.org/appengine v1.6.1 // indirect
	google.golang.org/genproto v0.0.0-20190707102051-710ae3a149df // indirect
	google.golang.org/grpc v1.22.0
)

replace github.com/grpc-ecosystem/grpc-gateway => github.com/moul/grpc-gateway v1.9.1-0.20190603230725-390f150e109c
