#!/bin/sh
protoc -I api api/karlchen.proto --go_out=plugins=grpc:api