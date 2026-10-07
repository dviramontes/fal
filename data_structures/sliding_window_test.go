package datastructures

import (
	"math"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		k        int
		expected []int
	}{
		{
			name:     "basic case",
			prices:   []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:        3,
			expected: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name:     "single element window",
			prices:   []int{1, 2, 3, 4, 5},
			k:        1,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "window equals array size",
			prices:   []int{5, 3, 4},
			k:        3,
			expected: []int{5},
		},
		{
			name:     "descending prices",
			prices:   []int{9, 8, 7, 6, 5},
			k:        2,
			expected: []int{9, 8, 7, 6},
		},
		{
			name:     "ascending prices",
			prices:   []int{1, 2, 3, 4, 5},
			k:        2,
			expected: []int{2, 3, 4, 5},
		},
		{
			name:     "all same values",
			prices:   []int{5, 5, 5, 5},
			k:        2,
			expected: []int{5, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxSlidingWindow(tt.prices, tt.k)
			if !equalIntSlice(result, tt.expected) {
				t.Errorf("MaxSlidingWindow(%v, %d) = %v, want %v",
					tt.prices, tt.k, result, tt.expected)
			}
		})
	}
}

func TestMovingAverage(t *testing.T) {
	ma := NewMovingAverage(3)
	if ma == nil {
		t.Fatal("NewMovingAverage returned nil")
	}

	tests := []struct {
		input    int
		expected float64
	}{
		{1, 1.0},
		{10, 5.5},
		{3, 4.666666666666667},
		{5, 6.0},
	}

	for i, tt := range tests {
		result := ma.Next(tt.input)
		if math.Abs(result-tt.expected) > 1e-9 {
			t.Errorf("Step %d: Next(%d) = %f, want %f",
				i+1, tt.input, result, tt.expected)
		}
	}
}

func equalIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
