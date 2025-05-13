#!/bin/bash
echo "PATH: $PATH"
which protoc
set -e

PROJECT_ROOT="$(dirname "$(dirname "$(readlink -f "$0")")")"
PROTO_ROOT="$PROJECT_ROOT/pkg/proto"
GO_OUT="$PROJECT_ROOT/pkg/proto/gen"

mkdir -p "$GO_OUT"

echo "Generating protobuf and gRPC code..."
protoc --proto_path="$PROTO_ROOT" \
  --go_out="$GO_OUT" --go_opt=paths=source_relative \
  --go-grpc_out="$GO_OUT" --go-grpc_opt=paths=source_relative \
  "$PROTO_ROOT"/*.proto

echo "Code generation completed."