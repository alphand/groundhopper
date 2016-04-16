FROM golang:1.6.1-alpine
MAINTAINER niko.darmawan@gmail.com

RUN apk add --update git && \
  addgroup app && \
  adduser -h /home/app -s /bin/false -G app -D app && \
  go get github.com/codegangsta/gin

ENV HOME=/home/app \
  GOPATH=/home/app/groundhopper/server \
  GOBIN=/home/app/groundhopper/server/bin \
  PORT=9000

USER app

WORKDIR $HOME/groundhopper/server/src
