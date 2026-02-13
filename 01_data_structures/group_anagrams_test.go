package datastructures

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// normalize sorts each group internally and sorts groups by length then first element
// so test assertions are deterministic regardless of return order.
func normalize(groups [][]string) [][]string {
	for _, g := range groups {
		sort.Strings(g)
	}
	sort.Slice(groups, func(i, j int) bool {
		if len(groups[i]) != len(groups[j]) {
			return len(groups[i]) < len(groups[j])
		}
		return groups[i][0] < groups[j][0]
	})
	return groups
}

func TestGroupAnagrams(t *testing.T) {
	t.Run("example case", func(t *testing.T) {
		strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
		expected := [][]string{{"hat"}, {"act", "cat"}, {"pots", "stop", "tops"}}
		assert.Equal(t, normalize(expected), normalize(GroupAnagrams(strs)))
	})

	t.Run("single string", func(t *testing.T) {
		strs := []string{"hello"}
		expected := [][]string{{"hello"}}
		assert.Equal(t, normalize(expected), normalize(GroupAnagrams(strs)))
	})

	t.Run("all anagrams", func(t *testing.T) {
		strs := []string{"abc", "bca", "cab"}
		expected := [][]string{{"abc", "bca", "cab"}}
		assert.Equal(t, normalize(expected), normalize(GroupAnagrams(strs)))
	})

	t.Run("no anagrams", func(t *testing.T) {
		strs := []string{"abc", "def", "ghi"}
		expected := [][]string{{"abc"}, {"def"}, {"ghi"}}
		assert.Equal(t, normalize(expected), normalize(GroupAnagrams(strs)))
	})

	t.Run("empty strings", func(t *testing.T) {
		strs := []string{"", ""}
		expected := [][]string{{"", ""}}
		assert.Equal(t, normalize(expected), normalize(GroupAnagrams(strs)))
	})

	t.Run("mix of empty and non-empty", func(t *testing.T) {
		strs := []string{"", "a", ""}
		expected := [][]string{{"a"}, {"", ""}}
		assert.Equal(t, normalize(expected), normalize(GroupAnagrams(strs)))
	})

	t.Run("single character strings", func(t *testing.T) {
		strs := []string{"a", "b", "a", "c", "b"}
		expected := [][]string{{"c"}, {"a", "a"}, {"b", "b"}}
		assert.Equal(t, normalize(expected), normalize(GroupAnagrams(strs)))
	})

	t.Run("duplicate strings", func(t *testing.T) {
		strs := []string{"eat", "eat", "eat"}
		expected := [][]string{{"eat", "eat", "eat"}}
		assert.Equal(t, normalize(expected), normalize(GroupAnagrams(strs)))
	})

	t.Run("longer words", func(t *testing.T) {
		strs := []string{"listen", "silent", "enlist", "hello", "world"}
		expected := [][]string{{"hello"}, {"world"}, {"enlist", "listen", "silent"}}
		assert.Equal(t, normalize(expected), normalize(GroupAnagrams(strs)))
	})

	t.Run("same letters different counts are not anagrams", func(t *testing.T) {
		strs := []string{"aab", "aba", "abb", "bba"}
		expected := [][]string{{"aab", "aba"}, {"abb", "bba"}}
		assert.Equal(t, normalize(expected), normalize(GroupAnagrams(strs)))
	})
}
