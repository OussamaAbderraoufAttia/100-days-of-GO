package main

import "testing"

func Sum(a, b int) int {
	return a + b
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2)
	}
}
