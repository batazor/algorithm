package matrix_search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchMatrix(t *testing.T) {
	in := []struct {
		name   string
		matrix [][]int
		target int
		got    bool
	}{
		{
			name: "case 1",
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 6,
			got:    true,
		},
		{
			name: "case 2",
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 20,
			got:    false,
		},
		{
			name: "case 3",
			matrix: [][]int{
				{-5},
			},
			target: -2,
			got:    false,
		},
		{
			name: "case 4",
			matrix: [][]int{
				{-1, 3},
			},
			target: 3,
			got:    true,
		},
		{
			name: "case 5",
			matrix: [][]int{
				{-1},
				{3},
				{5},
			},
			target: 5,
			got:    true,
		},
	}

	t.Run("N2", func(t *testing.T) {
		for _, v := range in {
			resp := SearchMatrix_N2(v.matrix, v.target)
			require.Equal(t, v.got, resp, v.name)
		}
	})

	t.Run("N_LOG_N", func(t *testing.T) {
		for _, v := range in {
			resp := SearchMatrix_N_LOG_N(v.matrix, v.target)
			require.Equal(t, v.got, resp, v.name)
		}
	})
}
