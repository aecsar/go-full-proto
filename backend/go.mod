module github.com/aecsar/go-proto

go 1.26.1

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.11-20260209202127-80ab13bee0bf.1
	connectrpc.com/connect v1.19.1
	google.golang.org/protobuf v1.36.11
)

tool (
	connectrpc.com/connect/cmd/protoc-gen-connect-go
	google.golang.org/protobuf/cmd/protoc-gen-go
)
