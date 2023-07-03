package unique_paths

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"3x7",
			args{
				m: 3,
				n: 7,
			},
			28,
		},
		{
			"3x2",
			args{
				m: 3,
				n: 2,
			},
			3,
		},
	}

	t.Run("2N_2M", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := UniquePaths_2N_2M(tt.args.m, tt.args.n); got != tt.want {
					require.Equal(t, tt.want, got)
				}
			})
		}
	})

	t.Run("NM", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := UniquePaths_NM(tt.args.m, tt.args.n); got != tt.want {
					require.Equal(t, tt.want, got)
				}
			})
		}
	})
}
