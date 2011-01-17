package main

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/keruch/go-tcp-pow/internal/server"
)

func main() {
	c, ok := os.LookupEnv("CHALLENGE_COMPLEXITY")
	if !ok {
		slog.Error("Can't find CHALLENGE_COMPLEXITY env var")
		return
	}
	complexity, err := strconv.Atoi(c)
	if err != nil {
		slog.With("err", err).Error("Invalid CHALLENGE_COMPLEXITY env")
		return
	}

	slog.Info("Starting TCP service")

	err = server.Run("0.0.0.0:9000", complexity)
	if err != nil {
		slog.With("err", err).Error("server.Run error")
	}
}
