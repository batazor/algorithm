package two_sum

import (
	"sort"
)

func TwoSum_N2(in []int, target int) []int {
	for i := 0; i < len(in); i++ {
		for j := i + 1; j < len(in); j++ {
			if in[i]+in[j] == target {
				return []int{in[i], in[j]}
			}
		}
	}

	return nil
}

func TwoSum_N_LOG_N(in []int, target int) []int {
	if ok := sort.IntsAreSorted(in); !ok {
		sort.Ints(in)
	}

	queue := map[int]struct{}{}
	for _, v := range in {
		if _, ok := queue[target-v]; ok {
			return []int{v, target - v}
		}

		queue[v] = struct{}{}
	}

	return nil
}

func TwoSum_N_LOG_N_BINARY_SEARCH(in []int, target int) []int {
	if ok := sort.IntsAreSorted(in); !ok {
		sort.Ints(in)
	}

	for i := 0; i < len(in); i++ {
		sum := target - in[i]
		left, right := i+1, len(in)-1
		for left <= right {
			mid := (left + right) / 2
			if in[mid] == sum {
				return []int{in[i], in[mid]}
			} else if in[mid] < sum {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return nil
}

func TwoSum_N(in []int, target int) []int {
	if ok := sort.IntsAreSorted(in); !ok {
		sort.Ints(in)
	}

	left, right := 0, len(in)-1
	for left < right {
		sum := in[left] + in[right]
		if sum == target {
			return []int{in[left], in[right]}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}
