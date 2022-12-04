ARG GO_VERSION=1.19.3

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY *.go ./
COPY . .

RUN go build -o /docker-ett

EXPOSE 9999

CMD [ "/docker-ett" ]