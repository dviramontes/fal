package datastructures

// Exercise 1: Sliding Window Maximum
// Given a stream of stock prices and a window size k, return the maximum price
// in each sliding window. This is common in financial analysis for finding
// local peaks in time series data.
//
// Example:
//   prices: [1, 3, -1, -3, 5, 3, 6, 7], k: 3
//   output: [3, 3, 5, 5, 6, 7]
//
// Hint: Use a deque (double-ended queue) for O(n) time complexity.
// A naive approach is O(n*k) - can you do better?

func MaxSlidingWindow(prices []int, k int) []int {
	// TODO: Implement this
	return nil
}

// Exercise 2: Moving Average from Data Stream
// Design a class to calculate the moving average of stock prices.
// This is fundamental for technical analysis indicators like SMA.
//
// Example:
//   ma := NewMovingAverage(3)
//   ma.Next(1) -> 1.0
//   ma.Next(10) -> 5.5
//   ma.Next(3) -> 4.666...
//   ma.Next(5) -> 6.0

type MovingAverage struct {
	// TODO: Add fields
}

func NewMovingAverage(size int) *MovingAverage {
	// TODO: Implement this
	return nil
}

func (ma *MovingAverage) Next(val int) float64 {
	// TODO: Implement this
	return 0
}
