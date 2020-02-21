package main

import "testing"

func BenchmarkTwoSum1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum1([]int{2, 7, 11, 15}, 9)
	}
}

func BenchmarkTwoSum2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum2([]int{2, 7, 11, 15}, 9)
	}
}
