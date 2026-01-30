package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeIntervals(t *testing.T) {
	t.Run("example case", func(t *testing.T) {
		intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
		expected := [][]int{{1, 6}, {8, 10}, {15, 18}}
		assert.Equal(t, expected, MergeIntervals(intervals))
	})

	t.Run("empty input", func(t *testing.T) {
		intervals := [][]int{}
		expected := [][]int{}
		assert.Equal(t, expected, MergeIntervals(intervals))
	})

	t.Run("single interval", func(t *testing.T) {
		intervals := [][]int{{1, 4}}
		expected := [][]int{{1, 4}}
		assert.Equal(t, expected, MergeIntervals(intervals))
	})

	t.Run("no overlaps", func(t *testing.T) {
		intervals := [][]int{{1, 2}, {4, 5}, {7, 8}}
		expected := [][]int{{1, 2}, {4, 5}, {7, 8}}
		assert.Equal(t, expected, MergeIntervals(intervals))
	})

	t.Run("all overlapping into one", func(t *testing.T) {
		intervals := [][]int{{1, 4}, {2, 5}, {3, 6}}
		expected := [][]int{{1, 6}}
		assert.Equal(t, expected, MergeIntervals(intervals))
	})

	t.Run("unsorted input", func(t *testing.T) {
		intervals := [][]int{{8, 10}, {1, 3}, {15, 18}, {2, 6}}
		expected := [][]int{{1, 6}, {8, 10}, {15, 18}}
		assert.Equal(t, expected, MergeIntervals(intervals))
	})

	t.Run("touching boundaries merge", func(t *testing.T) {
		intervals := [][]int{{1, 3}, {3, 5}, {5, 7}}
		expected := [][]int{{1, 7}}
		assert.Equal(t, expected, MergeIntervals(intervals))
	})

	t.Run("fully contained interval", func(t *testing.T) {
		intervals := [][]int{{1, 10}, {3, 5}, {6, 8}}
		expected := [][]int{{1, 10}}
		assert.Equal(t, expected, MergeIntervals(intervals))
	})

	t.Run("duplicate intervals", func(t *testing.T) {
		intervals := [][]int{{1, 4}, {1, 4}}
		expected := [][]int{{1, 4}}
		assert.Equal(t, expected, MergeIntervals(intervals))
	})

	t.Run("two intervals that merge", func(t *testing.T) {
		intervals := [][]int{{1, 4}, {4, 5}}
		expected := [][]int{{1, 5}}
		assert.Equal(t, expected, MergeIntervals(intervals))
	})

	t.Run("two intervals no overlap", func(t *testing.T) {
		intervals := [][]int{{1, 2}, {4, 5}}
		expected := [][]int{{1, 2}, {4, 5}}
		assert.Equal(t, expected, MergeIntervals(intervals))
	})

	t.Run("same start different end", func(t *testing.T) {
		intervals := [][]int{{1, 4}, {1, 6}}
		expected := [][]int{{1, 6}}
		assert.Equal(t, expected, MergeIntervals(intervals))
	})

	t.Run("multiple merge groups", func(t *testing.T) {
		intervals := [][]int{{1, 3}, {2, 4}, {6, 8}, {7, 9}, {11, 13}}
		expected := [][]int{{1, 4}, {6, 9}, {11, 13}}
		assert.Equal(t, expected, MergeIntervals(intervals))
	})

	t.Run("zero-length intervals", func(t *testing.T) {
		intervals := [][]int{{1, 1}, {2, 2}, {3, 3}}
		expected := [][]int{{1, 1}, {2, 2}, {3, 3}}
		assert.Equal(t, expected, MergeIntervals(intervals))
	})

	t.Run("zero-length interval absorbed", func(t *testing.T) {
		intervals := [][]int{{1, 5}, {3, 3}}
		expected := [][]int{{1, 5}}
		assert.Equal(t, expected, MergeIntervals(intervals))
	})
}
