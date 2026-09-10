package popcount

import "testing"

func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}

// BenchmarkTable - Exercise 2.3
func BenchmarkTable(b *testing.B) {
	bench(b, PopCount)
}

// BenchmarkTableLoop - Exercise 2.3
func BenchmarkTableLoop(b *testing.B) {
	bench(b, PopCountCycle)
}
