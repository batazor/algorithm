package unique_paths

func UniquePaths_2N_2M(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}

	if m == 1 || n == 1 {
		return 1
	}

	return UniquePaths_2N_2M(n-1, m) + UniquePaths_2N_2M(m-1, n)
}

func UniquePaths_NM(m int, n int) int {
	arr := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		arr[i] = make([]int, n+1)
	}

	return uniquePaths_NM_helper(m, n, arr)
}

func uniquePaths_NM_helper(m int, n int, arr [][]int) int {
	if m < 1 || n < 1 {
		return 0
	}

	if m == 1 || n == 1 {
		return 1
	}

	if arr[m][n] != 0 {
		return arr[m][n]
	}

	arr[m][n] = uniquePaths_NM_helper(m-1, n, arr) + uniquePaths_NM_helper(m, n-1, arr)
	return arr[m][n]
}
