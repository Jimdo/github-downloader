FROM golang:alpine

ADD . /go/src/github.com/Jimdo/github-downloader
WORKDIR /go/src/github.com/Jimdo/github-downloader

RUN apk --update add ca-certificates \
    && go install -v

WORKDIR /opt
ENTRYPOINT ["github-downloader"]
