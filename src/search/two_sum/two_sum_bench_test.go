package two_sum

import (
	"testing"
)

func generateTestData(n int) ([]int, int) {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i
	}
	target := n - 1
	return data, target
}

func BenchmarkTwoSum_N2(b *testing.B) {
	data, target := generateTestData(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum_N2(data, target)
	}
}

func BenchmarkTwoSum_N_LOG_N(b *testing.B) {
	data, target := generateTestData(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum_N_LOG_N(data, target)
	}
}

func BenchmarkTwoSum_N(b *testing.B) {
	data, target := generateTestData(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum_N(data, target)
	}
}
