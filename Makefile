install: grpcbin/grpcbin.pb.go
	go install -v .

grpcbin/grpcbin.pb.go: grpcbin/grpcbin.proto
	go generate

test:
	go test -v .

lint:
	gometalinter --disable-all --enable=errcheck --enable=vet --enable=vetshadow --enable=golint --enable=gas --enable=ineffassign --enable=goconst --enable=goimports --enable=gofmt --exclude="Binds to all network interfaces" --exclude="should have comment" --enable=staticcheck --enable=gosimple --enable=misspell --deadline=20s .

gentls:
	mkdir -p cert
	openssl genrsa -out cert/server.key 2048
	openssl ecparam -genkey -name secp384r1 -out cert/server.key
	openssl req -new -x509 -sha256 -key cert/server.key -out cert/server.crt -days 3650
