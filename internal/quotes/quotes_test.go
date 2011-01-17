package quotes_test

import (
	"testing"

	"github.com/keruch/go-tcp-pow/internal/quotes"
	"github.com/stretchr/testify/require"
)

func TestRandQuote(t *testing.T) {
	q := quotes.RandQuote()
	require.NotEmpty(t, q)
}
