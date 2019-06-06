module github.com/moul/grpcbin

go 1.12

require (
	github.com/grpc-ecosystem/grpc-gateway v0.0.0-20190514190838-b5f6fca70e91
	github.com/moul/pb v0.0.0-20180404114147-54bdd96e6a52
	github.com/rogpeppe/fastuuid v1.1.0 // indirect
	github.com/smartystreets/assertions v0.0.0-20190401211740-f487f9de1cd3 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a
	golang.org/x/crypto v0.0.0-20190605224628-f99c8df09eb5
	golang.org/x/net v0.0.0-20190603230018-60506f45cf65
	golang.org/x/sync v0.0.0-20190603230018-112230192c58 // indirect
	golang.org/x/sys v0.0.0-20190606135534-79a91cf218c4 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/appengine v1.6.0 // indirect
	google.golang.org/genproto v0.0.0-20190605224628-eb0b1bdb6ae6 // indirect
	google.golang.org/grpc v1.21.1
)

replace github.com/grpc-ecosystem/grpc-gateway => github.com/moul/grpc-gateway v1.9.1-0.20190603230725-390f150e109c
