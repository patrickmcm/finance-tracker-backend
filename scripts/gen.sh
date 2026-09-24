#!/bin/bash

cd ../api/v1 || exit
protoc --go_out=../../gen/api --go_opt=paths=source_relative --go-grpc_out=../../gen/api --go-grpc_opt=paths=source_relative instruments.proto

cd ../../ || exit
sqlc generate