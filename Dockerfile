FROM golang:1.20.0-alpine3.17 AS build

RUN apk add --update make

COPY . /app
WORKDIR /app

RUN go get
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/google/wire/cmd/wire@latest

RUN make build

FROM alpine:3.17

RUN mkdir /app
WORKDIR /app

COPY --from=build /app/bin/server.app /app

STOPSIGNAL SIGKILL
ENTRYPOINT [ "/app/server.app" ]
