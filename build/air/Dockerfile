FROM golang:1.21-alpine


RUN go version

WORKDIR /app-air

RUN go install github.com/cosmtrek/air@latest

COPY ../go.mod ./


RUN apk update && apk add --no-cache build-base

RUN go mod download

CMD ["air", "-c", "./build/air/.air.toml"]