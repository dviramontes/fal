package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountGreaterThanPreviousAverage(t *testing.T) {
	tests := []struct {
		name          string
		responseTimes []int32
		expected      int32
	}{
		{
			name:          "example case",
			responseTimes: []int32{100, 200, 150, 300},
			expected:      2,
		},
		{
			name:          "empty array",
			responseTimes: []int32{},
			expected:      0,
		},
		{
			name:          "single element",
			responseTimes: []int32{100},
			expected:      0,
		},
		{
			name:          "two elements - greater",
			responseTimes: []int32{100, 200},
			expected:      1,
		},
		{
			name:          "two elements - equal",
			responseTimes: []int32{100, 100},
			expected:      0,
		},
		{
			name:          "two elements - less",
			responseTimes: []int32{200, 100},
			expected:      0,
		},
		{
			name:          "all increasing",
			responseTimes: []int32{1, 2, 3, 4, 5},
			expected:      4, // all except first are greater than previous avg
		},
		{
			name:          "all decreasing",
			responseTimes: []int32{5, 4, 3, 2, 1},
			expected:      0,
		},
		{
			name:          "all same values",
			responseTimes: []int32{100, 100, 100, 100},
			expected:      0,
		},
		{
			name:          "large values - test int32 overflow handling",
			responseTimes: []int32{1000000000, 1000000000, 1000000001},
			expected:      1, // only last is greater than avg of first two
		},
		{
			name:          "spike pattern",
			responseTimes: []int32{10, 10, 10, 1000, 10},
			expected:      1, // only the spike (1000) is greater
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountGreaterThanPreviousAverage(tt.responseTimes)
			assert.Equal(t, tt.expected, result, "input: %v", tt.responseTimes)
		})
	}
}
