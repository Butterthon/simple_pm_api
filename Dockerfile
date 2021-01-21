FROM golang:1.15.6-alpine

ENV GO111MODULE on

ADD . /simple_pm_api

WORKDIR /simple_pm_api

RUN apk update \
    && apk add --no-cache git postgresql \
    && go build -o simple_pm_api

RUN go get github.com/cosmtrek/air \
    && go get github.com/swaggo/swag/cmd/swag \
    && go get github.com/lib/pq \
    && go get github.com/golang-migrate/migrate/v4
