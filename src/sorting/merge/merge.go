package merge

import (
	sort2 "sort"

	merge_two_sort_array "github.com/batazor/algorithm/src/leetCode/88"
)

func Sort(arr []int) []int {
	return merge(arr)
}

func merge(arr []int) []int {
	sort2.Sort(sort2.IntSlice(arr))

	if len(arr) > 2 {
		a := arr[:(len(arr)-1)/2]
		b := arr[(len(arr)-1)/2:]

		s1 := merge(a)
		s2 := merge(b)

		return sort(s1, s2)
	}

	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return []int{arr[1], arr[0]}
		}
		return arr
	}

	return arr
}

func sort(a []int, b []int) []int {
	return merge_two_sort_array.Merge(a, b)
}
