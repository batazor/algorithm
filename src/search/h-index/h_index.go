package h_index

import (
	"sort"
)

func GetHIndex_N2(citations []int) int {
	h := map[int]int{}
	cit_max := 0

	for c := range citations {
		for i := 0; i <= citations[c]; i++ {
			h[i]++
		}
	}

	for i := 0; i <= len(citations); i++ {
		if i <= h[i] {
			cit_max = i
		}
	}

	return cit_max
}

func GetHIndex_N_LOG_N(citations []int) int {
	sort.Sort(sort.IntSlice(citations))
	cit_max := 0

	for i := len(citations) - 1; i >= 0; i-- {
		if cit_max < citations[i] {
			cit_max++
		}
	}

	return cit_max
}

func GetHIndex_N(citations []int) int {
	h := make([]int, len(citations)+1)
	cit_max := 0

	for _, c := range citations {
		if c >= len(citations) {
			h[len(citations)]++
		} else {
			h[c]++
		}
	}

	new_citations := []int{}
	for k, v := range h {
		for i := 0; i < v; i++ {
			new_citations = append(new_citations, k)
		}
	}

	for i := len(new_citations) - 1; i >= 0; i-- {
		if cit_max < new_citations[i] {
			cit_max++
		}
	}

	return cit_max
}
