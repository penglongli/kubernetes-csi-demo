#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -v -x -o docker/fake-nfsserver main.go

cd docker/

docker build -t fake-nfsserver:0.1.0 .
