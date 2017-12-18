.PHONY: install
install: generate
	go install -v .

.PHONY: generate
generate: grpcbin/grpcbin.pb.go grpcbin/grpcbin.pb.gw.go

grpcbin/grpcbin.pb.go: grpcbin/grpcbin.proto
	protoc -I ./grpcbin -I./vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:./grpcbin ./grpcbin/grpcbin.proto

grpcbin/grpcbin.pb.gw.go: grpcbin/grpcbin.proto
	protoc -I ./grpcbin -I./vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=:./grpcbin ./grpcbin/grpcbin.proto

.PHONY: test
test:
	go test -v ./...

.PHONY: docker.test
docker.test:
	docker run -v $(PWD):/go/src/github.com/moul/grpcbin golang:1.8 go test -v github.com/moul/grpcbin/...

.PHONY: lint
lint:
	gometalinter --disable-all --enable=errcheck --enable=vet --enable=vetshadow --enable=golint --enable=gas --enable=ineffassign --enable=goconst --enable=goimports --enable=gofmt --exclude="Binds to all network interfaces" --exclude="should have comment" --enable=staticcheck --enable=gosimple --enable=misspell --deadline=20s .

gentls:
	mkdir -p cert
	openssl genrsa -out cert/server.key 2048
	openssl ecparam -genkey -name secp384r1 -out cert/server.key
	openssl req -new -x509 -sha256 -key cert/server.key -out cert/server.crt -days 3650
