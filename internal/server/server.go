package server

import (
	"bufio"
	"io"
	"log/slog"
	"net"

	"github.com/keruch/go-tcp-pow/internal/hashcash"
	"github.com/keruch/go-tcp-pow/internal/quotes"
)

func Run(addr string, complexity int) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, errX := listener.Accept()
		if errX != nil {
			return errX
		}

		go func() {
			errG := handleConnection(conn, complexity)
			if errG != nil && errG != io.EOF {
				slog.With("address", conn.RemoteAddr()).
					With("err", errG).
					Error("handleConnection error")
			}
		}()
	}
}

func handleConnection(conn net.Conn, complexity int) error {
	defer func() {
		slog.With("address", conn.RemoteAddr()).Info("User connection is closed")
		conn.Close()
	}()

	slog.With("address", conn.RemoteAddr()).Info("Accepted user connection")

	const incorrect = "incorrect challenge"
	var (
		reader = bufio.NewReader(conn)

		header       = hashcash.Generate(conn.RemoteAddr().String(), complexity)
		headerString = header.String()
	)

	_, err := conn.Write(wrapResp(headerString))
	if err != nil {
		return err
	}

	for {
		message, errX := reader.ReadString('\n')
		if errX != nil {
			return errX
		}
		message = message[:len(message)-1] // delete \n at the end

		switch {
		case hashcash.IsHashCorrect(headerString, message, header.ZeroBytes):
			_, errX = conn.Write(wrapResp(quotes.RandQuote()))
			if errX != nil {
				return errX
			}
		default:
			_, errX = conn.Write(wrapResp(incorrect))
			if errX != nil {
				return errX
			}
		}
	}
}

func wrapResp(resp string) []byte {
	return append([]byte(resp), '\n')
}
