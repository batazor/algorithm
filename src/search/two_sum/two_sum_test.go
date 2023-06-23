package two_sum

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum_N2(t *testing.T) {
	cases := []struct {
		in     []int
		target int
		want   []int
	}{
		{[]int{1, 2, 3, 4, 5}, 9, []int{4, 5}},
		{[]int{3, 5, -4, 8, 11, 1, -1, 6}, 10, []int{-1, 11}},
		{[]int{4, 6, 1}, 10, []int{4, 6}},
		{[]int{4, 6, 1}, 15, nil},
		{[]int{}, 6, nil},
	}

	for _, c := range cases {
		got := TwoSum_N2(c.in, c.target)
		sort.Ints(got)
		require.Equal(t, c.want, got)
	}
}

func TestTwoSum_N_LOG_N(t *testing.T) {
	cases := []struct {
		in     []int
		target int
		want   []int
	}{
		{[]int{1, 2, 3, 4, 5}, 9, []int{4, 5}},
		{[]int{3, 5, -4, 8, 11, 1, -1, 6}, 10, []int{-1, 11}},
		{[]int{4, 6, 1}, 10, []int{4, 6}},
		{[]int{4, 6, 1}, 15, nil},
		{[]int{}, 6, nil},
	}

	for _, c := range cases {
		got := TwoSum_N_LOG_N(c.in, c.target)
		sort.Ints(got)
		require.Equal(t, c.want, got)
	}
}

func TestTwoSum_N_LOG_N_BINARY_SEARCH(t *testing.T) {
	cases := []struct {
		in     []int
		target int
		want   []int
	}{
		{[]int{1, 2, 3, 4, 5}, 9, []int{4, 5}},
		{[]int{3, 5, -4, 8, 11, 1, -1, 6}, 10, []int{-1, 11}},
		{[]int{4, 6, 1}, 10, []int{4, 6}},
		{[]int{4, 6, 1}, 15, nil},
		{[]int{}, 6, nil},
	}

	for _, c := range cases {
		got := TwoSum_N_LOG_N_BINARY_SEARCH(c.in, c.target)
		sort.Ints(got)
		require.Equal(t, c.want, got)
	}
}

func TestTwoSum_N(t *testing.T) {
	cases := []struct {
		in     []int
		target int
		want   []int
	}{
		{[]int{1, 2, 3, 4, 5}, 9, []int{4, 5}},
		{[]int{3, 5, -4, 8, 11, 1, -1, 6}, 10, []int{-1, 11}},
		{[]int{4, 6, 1}, 10, []int{4, 6}},
		{[]int{4, 6, 1}, 15, nil},
		{[]int{}, 6, nil},
	}

	for _, c := range cases {
		got := TwoSum_N(c.in, c.target)
		sort.Ints(got)
		require.Equal(t, c.want, got)
	}
}
