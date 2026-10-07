package datastructures

// Exercise: Binary Search
// Given a sorted array of integers and a target value, return the index of the
// target if found. If not found, return -1.
//
// Constraints:
//   - The array is sorted in ascending order
//   - All elements are unique
//   - 1 <= len(nums) <= 10^4
//   - -10^4 <= nums[i], target <= 10^4
//
// Example 1:
//   nums: [-1, 0, 3, 5, 9, 12], target: 9
//   output: 4
//
// Example 2:
//   nums: [-1, 0, 3, 5, 9, 12], target: 2
//   output: -1
//
// Hint: Classic binary search - narrow the search space by half each iteration.
// Time: O(log n), Space: O(1)

func BinarySearch(nums []int, target int) int {
	return -1 // not found
}

// Exercise: Search Insert Position
// Given a sorted array and a target value, return the index where the target
// is found. If not found, return the index where it would be inserted to keep
// the array sorted.
//
// This is useful for maintaining sorted order when inserting into arrays,
// such as keeping a sorted list of price levels in an order book.
//
// Constraints:
//   - The array is sorted in ascending order
//   - All elements are unique
//   - 0 <= len(nums) <= 10^4
//
// Example 1:
//   nums: [1, 3, 5, 6], target: 5
//   output: 2 (found at index 2)
//
// Example 2:
//   nums: [1, 3, 5, 6], target: 2
//   output: 1 (would insert between 1 and 3)
//
// Example 3:
//   nums: [1, 3, 5, 6], target: 7
//   output: 4 (would insert at end)
//
// Example 4:
//   nums: [1, 3, 5, 6], target: 0
//   output: 0 (would insert at beginning)
//
// Hint: Similar to binary search, but consider what to return when target is not found.

func SearchInsertPosition(nums []int, target int) int {
	// TODO: Implement this
	return 0
}

// Exercise: Find First and Last Position
// Given a sorted array and a target value, find the starting and ending
// position of the target. If target is not found, return [-1, -1].
//
// This pattern is useful for finding ranges in sorted data, such as all trades
// within a price range or all events within a time window.
//
// Constraints:
//   - The array is sorted in ascending order
//   - 0 <= len(nums) <= 10^5
//   - Elements may have duplicates
//
// Example 1:
//   nums: [5, 7, 7, 8, 8, 10], target: 8
//   output: [3, 4]
//
// Example 2:
//   nums: [5, 7, 7, 8, 8, 10], target: 6
//   output: [-1, -1]
//
// Example 3:
//   nums: [], target: 0
//   output: [-1, -1]
//
// Hint: Run binary search twice - once to find the leftmost occurrence,
// once to find the rightmost. Time: O(log n), Space: O(1)

func SearchRange(nums []int, target int) []int {
	// TODO: Implement this
	return []int{-1, -1}
}
