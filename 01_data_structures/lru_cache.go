package datastructures

// Exercise 3: LRU Cache
// Design and implement an LRU (Least Recently Used) cache.
// This is critical for caching expensive computations in trading systems.
//
// Requirements:
// - Get(key) returns the value if it exists, otherwise -1
// - Put(key, value) inserts or updates the value
// - Both operations must be O(1) time complexity
// - When capacity is exceeded, evict the least recently used item
//
// Example:
//   cache := NewLRUCache(2)
//   cache.Put(1, 1)
//   cache.Put(2, 2)
//   cache.Get(1)    -> 1
//   cache.Put(3, 3) -> evicts key 2
//   cache.Get(2)    -> -1 (not found)
//
// Hint: Combine a hashmap with a doubly-linked list

type LRUCache struct {
	// TODO: Add fields
}

func NewLRUCache(capacity int) *LRUCache {
	// TODO: Implement this
	return nil
}

func (c *LRUCache) Get(key int) int {
	// TODO: Implement this
	return -1
}

func (c *LRUCache) Put(key, value int) {
	// TODO: Implement this
}
