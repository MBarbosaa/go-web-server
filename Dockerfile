# syntax=docker/dockerfile:1

FROM golang:1.18.1-alpine


RUN apk add git
WORKDIR /app
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /main

EXPOSE $PORT

CMD ["/main"]