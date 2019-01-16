FROM golang:stretch

LABEL maintainer="Michał Prył <mpryl@n-educatio.pl>" \
      description="Image for building Golang apps for Debian/Ubuntu Linux" \
      version="1.0"

RUN apt-get update && apt-get install -y --no-install-recommends git && \
    apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN useradd -ms /bin/bash -u 1000 docker \
    && chown -R docker.docker /home/docker

COPY entrypoint.sh /entrypoint.sh

USER docker

WORKDIR $GOPATH/src/app

CMD ["/entrypoint.sh"]
