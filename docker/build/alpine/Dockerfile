FROM golang:alpine

LABEL maintainer="Michał Prył <mpryl@n-educatio.pl>" \
      description="Image for building Golang apps for Alpine Linux" \
      version="1.0"

RUN apk add --no-cache git

RUN adduser -D -u 1000 docker \
    && chown -R docker.docker /home/docker

COPY entrypoint.sh /entrypoint.sh

USER docker

WORKDIR $GOPATH/src/app

CMD ["/entrypoint.sh"]
