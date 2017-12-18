# grpcbin
httpbin like for gRPC

## Demo server

http://grpcbin.m.42.am:9000 (insecure)

## Example

Example with [grpcc](https://github.com/njpatel/grpcc):

```console
$ wget -qN https://github.com/moul/grpcbin/raw/master/grpcbin/grpcbin.proto
$ npm install -g grpcc
$
$ grpcc -i -p ./grpcbin.proto --address grpcbin.m.42.am:9000
Connecting to grpcbin.GRPCBin on grpcbin.m.42.am:9000. Available globals:

  client - the client connection to GRPCBin
    index (EmptyMessage, callback) returns IndexReply
    dummyUnary (DummyMessage, callback) returns DummyMessage
    dummyServerStream (DummyMessage, callback) returns DummyMessage
    dummyClientStream (DummyMessage, callback) returns DummyMessage
    dummyBidirectionalStreamStream (DummyMessage, callback) returns DummyMessage

  printReply - function to easily print a unary call reply (alias: pr)
  streamReply - function to easily print stream call replies (alias: sr)
  createMetadata - convert JS objects into grpc metadata instances (alias: cm)

GRPCBin@grpcbin.m.42.am:9000> ^C
$
$ grpcc -i -p ./grpcbin.proto --address grpcbin.m.42.am:9000 --eval 'client.index({}, printReply)'
{
  "description": "gRPC testing server",
  "endpoints": [
    {
      "path": "index",
      "description": "This endpoint."
    },
    {
      "path": "dummyUnary",
      "description": "Unary endpoint that replies a received DummyMessage."
    },
    [...]
  ]
}
```



