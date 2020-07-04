GOPKG ?=	moul.io/grpcbin
DOCKER_IMAGE ?=	moul/grpcbin
GOBINS ?=	.

include rules.mk

.PHONY: gentls
gentls:
	mkdir -p cert
	openssl genrsa -out cert/server.key 2048
	openssl ecparam -genkey -name secp384r1 -out cert/server.key
	openssl req -new -x509 -sha256 -key cert/server.key -out cert/server.crt -days 3650

.PHONY: doc
doc:
	dot -Tsvg ./.assets/overview.dot > ./.assets/overview.svg
