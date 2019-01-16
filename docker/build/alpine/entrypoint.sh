#!/bin/sh
BUILD_DIR=/go/src/app/build/alpine

go get -d -v

mkdir -p $BUILD_DIR

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o "$BUILD_DIR/$APP_NAME"
