package datastructures

import (
	"math"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name     string
		tickers  []string
		k        int
		expected []string
	}{
		{
			name:     "basic case",
			tickers:  []string{"AAPL", "GOOGL", "AAPL", "MSFT", "AAPL", "GOOGL"},
			k:        2,
			expected: []string{"AAPL", "GOOGL"},
		},
		{
			name:     "k equals unique count",
			tickers:  []string{"AAPL", "GOOGL", "MSFT"},
			k:        3,
			expected: []string{"AAPL", "GOOGL", "MSFT"},
		},
		{
			name:     "single ticker",
			tickers:  []string{"AAPL", "AAPL", "AAPL"},
			k:        1,
			expected: []string{"AAPL"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TopKFrequent(tt.tickers, tt.k)
			if !sameElements(result, tt.expected) {
				t.Errorf("TopKFrequent(%v, %d) = %v, want %v (any order)",
					tt.tickers, tt.k, result, tt.expected)
			}
		})
	}
}

func TestMergeKSortedStreams(t *testing.T) {
	tests := []struct {
		name     string
		streams  [][]int
		expected []int
	}{
		{
			name:     "three streams",
			streams:  [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:     "empty streams",
			streams:  [][]int{},
			expected: []int{},
		},
		{
			name:     "single stream",
			streams:  [][]int{{1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "streams with empty",
			streams:  [][]int{{1, 2}, {}, {3, 4}},
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeKSortedStreams(tt.streams)
			if !equalIntSlice(result, tt.expected) {
				t.Errorf("MergeKSortedStreams(%v) = %v, want %v",
					tt.streams, result, tt.expected)
			}
		})
	}
}

func TestMedianFinder(t *testing.T) {
	mf := NewMedianFinder()
	if mf == nil {
		t.Fatal("NewMedianFinder returned nil")
	}

	tests := []struct {
		addNum   int
		expected float64
	}{
		{1, 1.0},
		{2, 1.5},
		{3, 2.0},
		{4, 2.5},
		{5, 3.0},
	}

	for i, tt := range tests {
		mf.AddNum(tt.addNum)
		result := mf.FindMedian()
		if math.Abs(result-tt.expected) > 1e-9 {
			t.Errorf("Step %d: after AddNum(%d), FindMedian() = %f, want %f",
				i+1, tt.addNum, result, tt.expected)
		}
	}
}

func sameElements(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sortedA := make([]string, len(a))
	sortedB := make([]string, len(b))
	copy(sortedA, a)
	copy(sortedB, b)
	sort.Strings(sortedA)
	sort.Strings(sortedB)
	for i := range sortedA {
		if sortedA[i] != sortedB[i] {
			return false
		}
	}
	return true
}
