package datastructures

// Exercise: Group Anagrams
// Given an array of strings, group the anagrams together. You can return the
// answer in any order. An anagram is a word formed by rearranging the letters
// of another word, using all the original letters exactly once.
//
// Example:
//   strs = ["act","pots","tops","cat","stop","hat"]
//   - "act" and "cat" are anagrams (both have letters a, c, t)
//   - "pots", "tops", and "stop" are anagrams (all have letters o, p, s, t)
//   - "hat" has no anagram partner
//   Result: [["hat"],["act","cat"],["stop","pots","tops"]]
//   (group order and order within groups may vary)
//
// Constraints:
//   - 1 <= len(strs) <= 10^4
//   - 0 <= len(strs[i]) <= 100
//   - strs[i] consists of lowercase English letters only
//
// Hint: Two strings are anagrams if they produce the same value when their
// characters are sorted. Use a map keyed by the sorted string to bucket
// anagrams together.
// Time: O(n * k log k) where n = len(strs), k = max string length
// Space: O(n * k)

func GroupAnagrams(strs []string) [][]string {
	// TODO: Implement this
	return nil
}
