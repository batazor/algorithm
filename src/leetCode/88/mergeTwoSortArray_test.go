package merge_two_sort_array

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"simple case",
			args{
				a: []int{1, 5, 6, 9},
				b: []int{2, 3, 40},
			},
			[]int{1, 2, 3, 5, 6, 9, 40},
		},
		{
			"simple case: reverse",
			args{
				a: []int{2, 3, 40},
				b: []int{1, 5, 6, 9},
			},
			[]int{1, 2, 3, 5, 6, 9, 40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				require.Equal(t, tt.want, got)
			}
		})
	}
}
