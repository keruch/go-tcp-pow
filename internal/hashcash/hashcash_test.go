package hashcash_test

import (
	"testing"

	"github.com/keruch/go-tcp-pow/internal/hashcash"
	"github.com/stretchr/testify/require"
)

const (
	addr  = "testaddr"
	bytes = 5
)

func TestDecode(t *testing.T) {
	h0 := hashcash.Generate(addr, bytes)

	h1, err := hashcash.Decode(h0.String())

	require.NoError(t, err)
	require.Equal(t, h0, h1)
}

func TestCompute(t *testing.T) {
	h0 := hashcash.Generate(addr, bytes)

	h1, ok := hashcash.Compute(h0)
	require.True(t, ok)

	hashcash.IsHashCorrect(h0.String(), h1.String(), bytes)
}
