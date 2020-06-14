#!/bin/sh
echo "generating go code ..."
protoc -I. api/karlchen.proto --go_out=plugins=grpc:api
echo "generating typescript code ..."
protoc -I api karlchen.proto --js_out=import_style=commonjs:frontend/src/api --grpc-web_out=import_style=typescript,mode=grpcwebtext:frontend/src/api
echo "done."