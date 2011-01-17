package main

import (
	"log/slog"
	"os"

	"github.com/keruch/go-tcp-pow/internal/client"
)

func main() {
	addr, ok := os.LookupEnv("SERVER_ADDR")
	if !ok {
		slog.Error("Can't find SERVER_ADDR env var")
		return
	}

	err := client.Run(addr)
	if err != nil {
		slog.With("err", err).Error("client.Run error")
	}
}
