#!/bin/sh
protoc -I. api/karlchen.proto --go_out=plugins=grpc:api