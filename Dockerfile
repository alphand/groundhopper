FROM golang:1.6.1-alpine
MAINTAINER niko.darmawan@gmail.com

RUN addgroup app &&\
  adduser -h /home/app -s /bin/false -G app -D app

ENV HOME=/home/app
# ENV GOPATH=/home/app/groundhopper/server/
# ENV GOBIN=/home/app/groundhopper/server/bin
USER app

WORKDIR $HOME/groundhopper/server/src
