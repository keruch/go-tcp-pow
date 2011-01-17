package hashcash

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Header represents a hashcash header
// https://en.wikipedia.org/wiki/Hashcash
type Header struct {
	Version   int    // Header format version
	ZeroBytes int    // Number of zero bytes
	Date      int64  // The time that the message was sent, UNIX ts
	Resource  string // An IP address (or email, optionally), encoded in base-64 format
	Random    string // String of random characters, encoded in base-64 format
	Counter   string // Counter, encoded in base-64 format
}

func (h Header) String() string {
	return fmt.Sprintf(
		"%d:%d:%d:%s::%s:%s",
		h.Version,
		h.ZeroBytes,
		h.Date,
		h.Resource,
		h.Random,
		h.Counter,
	)
}

func Generate(clientAddr string, zeroBytes int) Header {
	return Header{
		Version:   1,
		ZeroBytes: zeroBytes,
		Date:      time.Now().Unix(),
		Resource:  generateBase64String(clientAddr),
		Random:    generateBase64Int(rand.Uint64()),
		Counter:   "",
	}
}

func Decode(header string) (Header, error) {
	tokens := strings.Split(header, ":")

	version, err := strconv.Atoi(tokens[0])
	if err != nil {
		return Header{}, fmt.Errorf("can't decode header version: %w", err)
	}
	zeroBits, err := strconv.Atoi(tokens[1])
	if err != nil {
		return Header{}, fmt.Errorf("can't decode header zero bits: %w", err)
	}
	date, err := strconv.Atoi(tokens[2])
	if err != nil {
		return Header{}, fmt.Errorf("can't decode header date: %w", err)
	}

	return Header{
		Version:   version,
		ZeroBytes: zeroBits,
		Date:      int64(date),
		Resource:  tokens[3],
		// 4 is always empty here
		Random:  tokens[5],
		Counter: tokens[6],
	}, nil
}

func IsHashCorrect(old, new string, zeroCount int) bool {
	return strings.HasPrefix(new, old) && isCorrect(new, zeroCount)
}

func generateSHA1(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func generateBase64Int(input uint64) string {
	return generateBase64String(fmt.Sprintf("%d", input))
}

func generateBase64String(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

const zeroByte = 48

func isCorrect(header string, zeroCount int) bool {
	hash := generateSHA1(header)

	if zeroCount > len(hash) {
		return false
	}
	for _, ch := range hash[:zeroCount] {
		if ch != zeroByte {
			return false
		}
	}
	return true
}

func Compute(input Header) (Header, bool) {
	for i := uint64(0); i < math.MaxUint64; i++ {
		input.Counter = generateBase64Int(i)
		if isCorrect(input.String(), input.ZeroBytes) {
			return input, true
		}
	}
	return Header{}, false
}
