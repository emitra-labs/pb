# pb

All Protobuf schemas used across my Golang projects.

## Prerequisites

Install `protoc`. You can install the binary directly to ensure latest version by visiting the link below.

https://grpc.io/docs/protoc-installation/#install-pre-compiled-binaries-any-os

After that, install the Go plugins as follows:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/favadi/protoc-go-inject-tag@latest
```

## Quickstart

```bash
# Re-generate protobuf schemas
make gen
```
