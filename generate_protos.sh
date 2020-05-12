#!/bin/sh
protoc -I=. --go_out=. ./api/enums.proto --go_opt=paths=source_relative && \
protoc -I. api/karlchen.proto --go_out=plugins=grpc:api