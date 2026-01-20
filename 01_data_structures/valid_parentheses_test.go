package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidParentheses(t *testing.T) {
	t.Run("simple pair", func(t *testing.T) {
		assert.True(t, IsValidParentheses("()"))
	})

	t.Run("multiple different pairs", func(t *testing.T) {
		assert.True(t, IsValidParentheses("()[]{}"))
	})

	t.Run("mismatched brackets", func(t *testing.T) {
		assert.False(t, IsValidParentheses("(]"))
	})

	t.Run("nested brackets", func(t *testing.T) {
		assert.True(t, IsValidParentheses("([])"))
	})

	t.Run("deeply nested", func(t *testing.T) {
		assert.True(t, IsValidParentheses("{[()]}"))
	})

	t.Run("interleaved invalid", func(t *testing.T) {
		assert.False(t, IsValidParentheses("([)]"))
	})

	t.Run("single open bracket", func(t *testing.T) {
		assert.False(t, IsValidParentheses("("))
	})

	t.Run("single close bracket", func(t *testing.T) {
		assert.False(t, IsValidParentheses(")"))
	})

	t.Run("extra closing bracket", func(t *testing.T) {
		assert.False(t, IsValidParentheses("())"))
	})

	t.Run("extra opening bracket", func(t *testing.T) {
		assert.False(t, IsValidParentheses("(()"))
	})

	t.Run("empty string", func(t *testing.T) {
		assert.True(t, IsValidParentheses(""))
	})

	t.Run("long valid sequence", func(t *testing.T) {
		assert.True(t, IsValidParentheses("(((())))[[[]]]{{{}}}"))
	})

	t.Run("complex nested valid", func(t *testing.T) {
		assert.True(t, IsValidParentheses("{[]([{}])}"))
	})

	t.Run("wrong closing order", func(t *testing.T) {
		assert.False(t, IsValidParentheses("{[}]"))
	})
}
