package h_index

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetHIndex(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Case 1", []int{3, 0, 6, 1, 5}, 3},
		{"Case 2", []int{1, 1, 1, 1}, 1},
		{"Case 3", []int{5, 5, 5, 5}, 4},
		{"Case 4", []int{1, 4, 1, 4, 2, 1, 2, 3, 5, 6}, 4},
		{"Case 5", []int{25, 8, 5, 3, 3}, 3},
		{"Case 6", []int{}, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name+"_N2", func(t *testing.T) {
			out := GetHIndex_N2(tc.input)
			require.Equal(t, tc.expected, out)
		})

		t.Run(tc.name+"_NlogN", func(t *testing.T) {
			out := GetHIndex_N_LOG_N(tc.input)
			require.Equal(t, tc.expected, out)
		})

		t.Run(tc.name+"_N", func(t *testing.T) {
			out := GetHIndex_N(tc.input)
			require.Equal(t, tc.expected, out)
		})
	}
}
