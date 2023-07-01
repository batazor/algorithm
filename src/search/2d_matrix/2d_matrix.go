package matrix_search

func SearchMatrix_N2(in [][]int, target int) bool {
	if len(in) == 0 {
		return false
	}

	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[i]); j++ {
			if in[i][j] == target {
				return true
			}
		}
	}

	return false
}

func SearchMatrix_N_LOG_N(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	w := 0
	h := len(matrix[0]) - 1

	for w < len(matrix) && h >= 0 {
		if matrix[w][h] == target {
			return true
		}

		if matrix[w][h] > target {
			h--
		} else {
			w++
		}
	}

	return false
}
