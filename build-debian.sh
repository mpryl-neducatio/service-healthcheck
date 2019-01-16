#!/bin/bash

docker build -t "go-builder:debian" -f "docker/build/debian/Dockerfile" ./docker/build/debian

docker run --rm -e APP_NAME="healthcheck" -v $(pwd):/go/src/app go-builder:debian
