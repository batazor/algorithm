package merge

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sort",
			args: args{arr: []int{3, 7, 1, 9, 2}},
			want: []int{1, 2, 3, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				require.Equal(t, tt.want, got)
			}
		})
	}
}
