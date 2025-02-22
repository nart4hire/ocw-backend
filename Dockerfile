FROM golang:1.20.0-alpine3.17 AS build

RUN apk add --update make
RUN apk add --update git

COPY . /app
WORKDIR /app

RUN go get
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.10
RUN go install github.com/google/wire/cmd/wire@latest

RUN make build

FROM alpine:3.17

RUN mkdir /app
WORKDIR /app

COPY --from=build /app/bin/server.app /app
RUN touch /app/.env
RUN touch /app/.env.local

STOPSIGNAL SIGKILL
ENTRYPOINT [ "/app/server.app" ]
