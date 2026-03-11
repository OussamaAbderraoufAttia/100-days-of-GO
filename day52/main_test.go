package main

import "testing"

func Sum(a, b int) int {
	return a + b
}

func TestSum(t *testing.T) {
	// Our test table
	tests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
		{-5, -5, -10},
	}

	for _, tc := range tests {
		result := Sum(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Sum(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
		}
	}
}
