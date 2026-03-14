# Go Full Proto

A full stack app with go backend and svelte frontend, shared protobuf schema and connectrpc, using [buf](https://buf.build) for protobuf schema management.

## Getting Started

### Prerequisites

- [Buf](https://buf.build/docs/installation)

### Running the app

1. Run `buf generate` to generate the protobuf code
2. Run `go run ./cmd/server` to start the server
