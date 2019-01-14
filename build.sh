#!/usr/bin/env bash

set -exo pipefail

docker run -it --rm -v $PWD:/go/src/github.com/mblink/drone golang:1.11 /bin/sh -c '
  cd /go/src/github.com/mblink/drone && \
  go get -u github.com/drone/drone-ui/dist && \
  go get -u golang.org/x/tools/cmd/cover && \
  go get -u golang.org/x/net/context && \
  go get -u golang.org/x/net/context/ctxhttp && \
  go get -u github.com/golang/protobuf/proto && \
  go get -u github.com/golang/protobuf/protoc-gen-go && \
  sh .drone.sh'
