FROM golang:1.21.5-alpine3.18 as builder

WORKDIR /build

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o client ./cmd/client

FROM alpine:3.18

WORKDIR /root/

COPY --from=builder /build/client .

CMD ["./client"]
