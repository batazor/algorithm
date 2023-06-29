package sum_from_array

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumFromArray(t *testing.T) {
	in := []int{7, 2, -5, 1, 1, -1, 5, -5}
	target := 5
	want := [][]int{
		{7, 2, -5, 1},
		{7, 2, -5, 1, 1, -1},
		{7, 2, -5, 1, 1, -1, 5, -5},
		{1, -1, 5},
		{5},
	}

	t.Run("N2", func(t *testing.T) {
		got := SumFromArray_N2(in, target)
		require.Equal(t, want, got)
	})

	t.Run("HASH", func(t *testing.T) {
		got := SumFromArray_HASH(in, target)
		require.Len(t, got, 5)
	})
}
