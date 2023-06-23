package A5B3C4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompress(t *testing.T) {
	in := "AAAAABBBCCCAAA"
	expected := "A5B3C3A3"

	result := Compress(in)
	require.Equal(t, expected, result)
}
