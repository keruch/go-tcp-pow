FROM golang:1.21.5-alpine3.18 as builder

WORKDIR /build

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o server ./cmd/server/

FROM alpine:3.18

WORKDIR /root/

COPY --from=builder /build/server .

EXPOSE 9000

CMD ["./server"]
