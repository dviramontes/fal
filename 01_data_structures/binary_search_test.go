package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	t.Run("target found in middle", func(t *testing.T) {
		nums := []int{-1, 0, 3, 5, 9, 12}
		assert.Equal(t, 4, BinarySearch(nums, 9))
	})
	//
	t.Run("target not found", func(t *testing.T) {
		nums := []int{-1, 0, 3, 5, 9, 12}
		assert.Equal(t, -1, BinarySearch(nums, 2))
	})
	//
	t.Run("target at beginning", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		assert.Equal(t, 0, BinarySearch(nums, 1))
	})

	t.Run("target at end", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		assert.Equal(t, 4, BinarySearch(nums, 5))
	})

	t.Run("single element found", func(t *testing.T) {
		nums := []int{5}
		assert.Equal(t, 0, BinarySearch(nums, 5))
	})

	t.Run("single element not found", func(t *testing.T) {
		nums := []int{5}
		assert.Equal(t, -1, BinarySearch(nums, 3))
	})

	t.Run("two elements first", func(t *testing.T) {
		nums := []int{1, 3}
		assert.Equal(t, 0, BinarySearch(nums, 1))
	})

	t.Run("two elements second", func(t *testing.T) {
		nums := []int{1, 3}
		assert.Equal(t, 1, BinarySearch(nums, 3))
	})

	t.Run("negative numbers", func(t *testing.T) {
		nums := []int{-10, -5, -3, -1, 0}
		assert.Equal(t, 1, BinarySearch(nums, -5))
	})

	t.Run("target smaller than all", func(t *testing.T) {
		nums := []int{2, 4, 6, 8}
		assert.Equal(t, -1, BinarySearch(nums, 1))
	})

	t.Run("target larger than all", func(t *testing.T) {
		nums := []int{2, 4, 6, 8}
		assert.Equal(t, -1, BinarySearch(nums, 10))
	})
}

//func TestSearchInsertPosition(t *testing.T) {
//	t.Run("target found", func(t *testing.T) {
//		nums := []int{1, 3, 5, 6}
//		assert.Equal(t, 2, SearchInsertPosition(nums, 5))
//	})
//
//	t.Run("insert in middle", func(t *testing.T) {
//		nums := []int{1, 3, 5, 6}
//		assert.Equal(t, 1, SearchInsertPosition(nums, 2))
//	})
//
//	t.Run("insert at end", func(t *testing.T) {
//		nums := []int{1, 3, 5, 6}
//		assert.Equal(t, 4, SearchInsertPosition(nums, 7))
//	})
//
//	t.Run("insert at beginning", func(t *testing.T) {
//		nums := []int{1, 3, 5, 6}
//		assert.Equal(t, 0, SearchInsertPosition(nums, 0))
//	})
//
//	t.Run("empty array", func(t *testing.T) {
//		nums := []int{}
//		assert.Equal(t, 0, SearchInsertPosition(nums, 5))
//	})
//
//	t.Run("single element insert before", func(t *testing.T) {
//		nums := []int{3}
//		assert.Equal(t, 0, SearchInsertPosition(nums, 1))
//	})
//
//	t.Run("single element insert after", func(t *testing.T) {
//		nums := []int{3}
//		assert.Equal(t, 1, SearchInsertPosition(nums, 5))
//	})
//
//	t.Run("single element found", func(t *testing.T) {
//		nums := []int{3}
//		assert.Equal(t, 0, SearchInsertPosition(nums, 3))
//	})
//
//	t.Run("insert between consecutive", func(t *testing.T) {
//		nums := []int{1, 2, 4, 5}
//		assert.Equal(t, 2, SearchInsertPosition(nums, 3))
//	})
//}

func TestSearchRange(t *testing.T) {
	t.Run("multiple occurrences in middle", func(t *testing.T) {
		nums := []int{5, 7, 7, 8, 8, 10}
		assert.Equal(t, []int{3, 4}, SearchRange(nums, 8))
	})

	t.Run("target not found", func(t *testing.T) {
		nums := []int{5, 7, 7, 8, 8, 10}
		assert.Equal(t, []int{-1, -1}, SearchRange(nums, 6))
	})

	t.Run("empty array", func(t *testing.T) {
		nums := []int{}
		assert.Equal(t, []int{-1, -1}, SearchRange(nums, 0))
	})

	t.Run("single occurrence", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		assert.Equal(t, []int{2, 2}, SearchRange(nums, 3))
	})

	t.Run("all same elements", func(t *testing.T) {
		nums := []int{5, 5, 5, 5, 5}
		assert.Equal(t, []int{0, 4}, SearchRange(nums, 5))
	})

	t.Run("range at beginning", func(t *testing.T) {
		nums := []int{1, 1, 1, 2, 3, 4}
		assert.Equal(t, []int{0, 2}, SearchRange(nums, 1))
	})

	t.Run("range at end", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 4, 4}
		assert.Equal(t, []int{3, 5}, SearchRange(nums, 4))
	})

	t.Run("single element found", func(t *testing.T) {
		nums := []int{1}
		assert.Equal(t, []int{0, 0}, SearchRange(nums, 1))
	})

	t.Run("single element not found", func(t *testing.T) {
		nums := []int{1}
		assert.Equal(t, []int{-1, -1}, SearchRange(nums, 2))
	})

	t.Run("target smaller than all", func(t *testing.T) {
		nums := []int{2, 2, 3, 3}
		assert.Equal(t, []int{-1, -1}, SearchRange(nums, 1))
	})

	t.Run("target larger than all", func(t *testing.T) {
		nums := []int{2, 2, 3, 3}
		assert.Equal(t, []int{-1, -1}, SearchRange(nums, 4))
	})

	t.Run("two elements both same", func(t *testing.T) {
		nums := []int{2, 2}
		assert.Equal(t, []int{0, 1}, SearchRange(nums, 2))
	})
}
