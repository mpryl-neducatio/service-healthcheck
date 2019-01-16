FROM golang:1.11-stretch

MAINTAINER Michał Prył <mpryl@n-educatio.pl>

ENV TERMINFO=/opt/share/terminfo \
    TERM=xterm \
    DEBIAN_FRONTEND=noninteractive \
    GOPATH=/project/packages

RUN apt-get update && apt-get install -y \
  curl \
  htop \
  nano \
  git && \
  apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

COPY docker/.bashrc /root/.bashrc

RUN useradd -ms /bin/bash docker \
    && chown -R docker.docker /home/docker

COPY docker/.bashrc /home/docker/.bashrc

USER docker

WORKDIR /project
