package client

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"time"

	"github.com/keruch/go-tcp-pow/internal/hashcash"
)

func Run(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	slog.With("address", conn.RemoteAddr()).Info("Connected to the server")

	reader := bufio.NewReader(conn)

	challengeString, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	challenge, err := hashcash.Decode(challengeString)
	if err != nil {
		return err
	}

	header, ok := hashcash.Compute(challenge)
	if !ok {
		return fmt.Errorf("can't compute challenge")
	}
	headerString := header.String()

	for {
		_, errX := conn.Write(wrapResp(headerString))
		if errX != nil {
			return errX
		}

		resp, errX := reader.ReadString('\n')
		if errX != nil {
			return errX
		}

		fmt.Print("Quote:", resp)

		time.Sleep(3 * time.Second)
	}
}

func wrapResp(resp string) []byte {
	return append([]byte(resp), '\n')
}
