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
