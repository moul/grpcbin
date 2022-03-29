# dynamic config
ARG             BUILD_DATE
ARG             VCS_REF
ARG             VERSION

# build
FROM            golang:1.18.0-alpine as builder
RUN             apk add --no-cache git gcc musl-dev make
ENV             GO111MODULE=on
WORKDIR         /go/src/moul.io/grpcbin
COPY            go.* ./
RUN             go mod download
COPY            . ./
#RUN             make install
RUN             go build -o /go/bin/grpcbin -ldflags "-extldflags \"-static\"" -v

# minimalist runtime
FROM alpine:3.15.3
LABEL           org.label-schema.build-date=$BUILD_DATE \
                org.label-schema.name="grpcbin" \
                org.label-schema.description="" \
                org.label-schema.url="https://moul.io/grpcbin/" \
                org.label-schema.vcs-ref=$VCS_REF \
                org.label-schema.vcs-url="https://github.com/moul/grpcbin" \
                org.label-schema.vendor="Manfred Touron" \
                org.label-schema.version=$VERSION \
                org.label-schema.schema-version="1.0" \
                org.label-schema.cmd="docker run -i -t --rm moul/grpcbin" \
                org.label-schema.help="docker exec -it $CONTAINER grpcbin --help"
RUN             apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY            --from=builder /go/bin/grpcbin /bin/grpcbin
COPY            --from=builder /go/src/moul.io/grpcbin/cert /root/cert
WORKDIR         /root
EXPOSE          9000 9001 80
ENTRYPOINT      ["/bin/grpcbin"]