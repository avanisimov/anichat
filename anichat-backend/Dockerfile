FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN apk add --no-cache ca-certificates git && \
    go mod download

COPY . .

RUN go build -o /usr/local/bin/anichat ./cmd/server

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/anichat"]
