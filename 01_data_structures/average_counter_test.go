package datastructures

import "testing"

func TestCountGreaterThanPreviousAverage(t *testing.T) {
	tests := []struct {
		name          string
		responseTimes []int
		expected      int
	}{
		{
			name:          "example case",
			responseTimes: []int{100, 200, 150, 300},
			expected:      2,
		},
		{
			name:          "empty array",
			responseTimes: []int{},
			expected:      0,
		},
		{
			name:          "single element",
			responseTimes: []int{100},
			expected:      0,
		},
		{
			name:          "two elements - greater",
			responseTimes: []int{100, 200},
			expected:      1,
		},
		{
			name:          "two elements - equal",
			responseTimes: []int{100, 100},
			expected:      0,
		},
		{
			name:          "two elements - less",
			responseTimes: []int{200, 100},
			expected:      0,
		},
		{
			name:          "all increasing",
			responseTimes: []int{1, 2, 3, 4, 5},
			expected:      4, // all except first are greater than previous avg
		},
		{
			name:          "all decreasing",
			responseTimes: []int{5, 4, 3, 2, 1},
			expected:      0,
		},
		{
			name:          "all same values",
			responseTimes: []int{100, 100, 100, 100},
			expected:      0,
		},
		{
			name:          "large values",
			responseTimes: []int{1000000000, 1000000000, 1000000001},
			expected:      1, // only last is greater than avg of first two
		},
		{
			name:          "spike pattern",
			responseTimes: []int{10, 10, 10, 1000, 10},
			expected:      1, // only the spike (1000) is greater
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountGreaterThanPreviousAverage(tt.responseTimes)
			if result != tt.expected {
				t.Errorf("CountGreaterThanPreviousAverage(%v) = %d, want %d",
					tt.responseTimes, result, tt.expected)
			}
		})
	}
}
