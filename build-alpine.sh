#!/bin/bash

docker build -t "go-builder:alpine" -f "docker/build/alpine/Dockerfile" ./docker/build/alpine

docker run --rm -e APP_NAME="healthcheck" -v $(pwd):/go/src/app go-builder:alpine
