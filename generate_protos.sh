#!/bin/sh
echo "generating go code ..."
# see https://grpc.io/docs/languages/go/quickstart/#prerequisites
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/karlchen.proto
echo "generating typescript code ..."
protoc -I api karlchen.proto --js_out=import_style=commonjs:frontend/src/api --grpc-web_out=import_style=typescript,mode=grpcwebtext:frontend/src/api
echo "done."
