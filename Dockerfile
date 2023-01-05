FROM golang:1.19-alpine

WORKDIR /app

RUN apk update && apk upgrade && apk add --update alpine-sdk && \
    apk add --no-cache 

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

ENTRYPOINT [ "go", "run", "main.go" ]