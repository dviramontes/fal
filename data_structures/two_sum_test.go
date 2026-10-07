package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	t.Run("example 1: basic case", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		target := 9
		result := TwoSum(nums, target)

		if !assert.Len(t, result, 2) {
			return
		}
		assert.Equal(t, target, nums[result[0]]+nums[result[1]])
	})

	t.Run("example 2: non-adjacent elements", func(t *testing.T) {
		nums := []int{3, 2, 4}
		target := 6
		result := TwoSum(nums, target)

		if !assert.Len(t, result, 2) {
			return
		}
		assert.Equal(t, target, nums[result[0]]+nums[result[1]])
	})

	t.Run("example 3: duplicate values", func(t *testing.T) {
		nums := []int{3, 3}
		target := 6
		result := TwoSum(nums, target)

		if !assert.Len(t, result, 2) {
			return
		}
		assert.NotEqual(t, result[0], result[1], "cannot use same element twice")
		assert.Equal(t, target, nums[result[0]]+nums[result[1]])
	})

	t.Run("negative numbers", func(t *testing.T) {
		nums := []int{-1, -2, -3, -4, -5}
		target := -8
		result := TwoSum(nums, target)

		if !assert.Len(t, result, 2) {
			return
		}
		assert.Equal(t, target, nums[result[0]]+nums[result[1]])
	})

	t.Run("mixed positive and negative", func(t *testing.T) {
		nums := []int{-3, 4, 3, 90}
		target := 0
		result := TwoSum(nums, target)

		if !assert.Len(t, result, 2) {
			return
		}
		assert.Equal(t, target, nums[result[0]]+nums[result[1]])
	})

	t.Run("large numbers", func(t *testing.T) {
		nums := []int{1000000000, -1000000000, 3, 4}
		target := 0
		result := TwoSum(nums, target)

		if !assert.Len(t, result, 2) {
			return
		}
		assert.Equal(t, target, nums[result[0]]+nums[result[1]])
	})

	t.Run("answer at end of array", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6}
		target := 11
		result := TwoSum(nums, target)

		if !assert.Len(t, result, 2) {
			return
		}
		assert.Equal(t, target, nums[result[0]]+nums[result[1]])
	})

	t.Run("minimum length array", func(t *testing.T) {
		nums := []int{1, 2}
		target := 3
		result := TwoSum(nums, target)

		if !assert.Len(t, result, 2) {
			return
		}
		assert.Equal(t, target, nums[result[0]]+nums[result[1]])
	})
}
