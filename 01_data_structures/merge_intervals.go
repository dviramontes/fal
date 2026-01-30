package datastructures

// Exercise: Merge and Sort Intervals
// Given an array of intervals [startTime, endTime], merge all overlapping
// intervals and return a sorted array of non-overlapping intervals.
//
// Two intervals overlap if one starts before or when the other ends.
// After merging, intervals should be sorted by start time.
//
// Example:
//   intervals = [[1,3],[2,6],[8,10],[15,18]]
//   - Sort by start time (already sorted).
//   - Initialize merged with [1,3].
//   - [2,6] overlaps [1,3] (2 <= 3), merge into [1,6].
//   - [8,10] does not overlap [1,6] (8 > 6), append.
//   - [15,18] does not overlap [8,10] (15 > 10), append.
//   Result: [[1,6],[8,10],[15,18]]
//
// Constraints:
//   - 0 <= len(intervals) <= 10^4
//   - intervals[i].length == 2
//   - 0 <= startTime <= endTime <= 10^4
//   - Intervals may be in any order
//   - Duplicates and fully contained intervals are allowed
//
// Hint: Sort intervals by start time first, then iterate and merge
// overlapping intervals by comparing current start with last merged end.
// Time: O(n log n), Space: O(n)

func MergeIntervals(intervals [][]int) [][]int {
	// TODO: Implement this
	return nil
}
