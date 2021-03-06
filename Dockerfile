FROM golang:1.16-alpine3.13

ARG RELEASE_VERSION
ARG RELEASE_COMMIT
ARG RELEASE_BUILD_TIME

ADD . /app
WORKDIR /app

RUN go mod download
RUN export PROJECT=$(head -n1 go.mod | sed -r 's/module (.*)/\1/') && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s -X $PROJECT/internal/version.Version=$RELEASE_VERSION \
        -X ${PROJECT}/internal/version.Commit=$RELEASE_COMMIT \
        -X ${PROJECT}/internal/version.BuildTime=$RELEASE_BUILD_TIME" \
    -o /gorelease /app/cmd/*
