package datastructures

// Exercise: Count Elements Greater Than Previous Average
// Given an array of positive integers, return the number of elements that are
// strictly greater than the average of all previous elements. Skip the first element.
//
// This is useful for detecting anomalies in time series data, such as
// response times that spike above historical averages.
//
// Example:
//   responseTimes = [100, 200, 150, 300]
//   - Day 0: 100 (no previous days, skip)
//   - Day 1: 200 > average(100) = 100 → count = 1
//   - Day 2: 150 vs average(100, 200) = 150 → not greater → count = 1
//   - Day 3: 300 > average(100, 200, 150) = 150 → count = 2
//   Return 2
//
// Constraints:
//   0 <= responseTimes.length <= 1000
//   1 <= responseTimes[i] <= 10^9
//
// Note: Uses int32 to match HackerRank problem format.
// Be careful with integer overflow when summing - you may need int64 for the running sum.
//
// Hint: Track running sum to calculate average in O(1) per element

func CountGreaterThanPreviousAverage(responseTimes []int32) int32 {
	var count int32
	// TODO
	return count
}
