package merge_two_sort_array

func Merge(a []int, b []int) []int {
	m := len(a)
	n := len(b)

	res := make([]int, m+n, m+n)

	for m >= 1 && n >= 1 {
		if a[m-1] > b[n-1] {
			res[m+n-1] = a[m-1]
			m--
		} else {
			res[m+n-1] = b[n-1]
			n--
		}
	}

	if m > 0 {
		for i := 0; i < m; i++ {
			res[i] = a[i]
		}
	}

	if n > 0 {
		for i := 0; i < n; i++ {
			res[i] = b[i]
		}
	}

	return res
}
