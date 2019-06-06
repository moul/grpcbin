GOENV ?= GO111MODULE=on
GO ?= $(GOENV) go

install:
	$(GO) install -v .

test:
	$(GO) test -v ./...

download:
	$(GO) mod download

docker.test:
	docker run -e "$(GOENV)" -v $(PWD):/go/src/github.com/moul/grpcbin golang:1.11 go test -v github.com/moul/grpcbin/...

install-dev:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b `go env GOPATH`/bin v1.16.0

lint:
	golangci-lint run --no-config --deadline=60s --disable-all --enable=errcheck --enable=vet --enable=vetshadow --enable=golint --enable=ineffassign --enable=goconst --enable=goimports --enable=gofmt --exclude="Binds to all network interfaces" --exclude="should have comment" --enable=staticcheck --enable=gosimple --enable=misspell . ./handler/..

gentls:
	mkdir -p cert
	openssl genrsa -out cert/server.key 2048
	openssl ecparam -genkey -name secp384r1 -out cert/server.key
	openssl req -new -x509 -sha256 -key cert/server.key -out cert/server.crt -days 3650

doc:
	dot -Tsvg ./.assets/overview.dot > ./.assets/overview.svg
