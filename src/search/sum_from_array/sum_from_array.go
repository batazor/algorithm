package sum_from_array

func SumFromArray_N2(in []int, target int) [][]int {
	result := [][]int{}

	for left := 0; left < len(in); left++ {
		sum := 0

		for right := left; right < len(in); right++ {
			sum += in[right]
			if sum == target {
				result = append(result, in[left:right+1])
			}
		}
	}

	return result
}

func SumFromArray_HASH(in []int, target int) [][]int {
	result := [][]int{}

	h := map[int][]int{}
	h[0] = []int{-1} // add -1 to denote the start of array

	sum := 0
	for i := 0; i < len(in); i++ {
		sum += in[i]
		h[sum] = append(h[sum], i) // append current index to the sum

		if indices, ok := h[sum-target]; ok {
			for _, index := range indices {
				result = append(result, append([]int(nil), in[index+1:i+1]...))
			}
		}
	}

	return result
}
