FROM golang:1.21-alpine

WORKDIR /app

RUN apk update && apk add --no-cache build-base

COPY ../go.mod ./

RUN go mod download

CMD ["go","run","cmd/server/main.go"]