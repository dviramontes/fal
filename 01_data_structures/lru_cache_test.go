package datastructures

import "testing"

func TestLRUCache(t *testing.T) {
	t.Run("basic operations", func(t *testing.T) {
		cache := NewLRUCache(2)
		if cache == nil {
			t.Fatal("NewLRUCache returned nil")
		}

		cache.Put(1, 1)
		cache.Put(2, 2)

		if got := cache.Get(1); got != 1 {
			t.Errorf("Get(1) = %d, want 1", got)
		}

		// This should evict key 2
		cache.Put(3, 3)

		if got := cache.Get(2); got != -1 {
			t.Errorf("Get(2) = %d, want -1 (evicted)", got)
		}

		if got := cache.Get(3); got != 3 {
			t.Errorf("Get(3) = %d, want 3", got)
		}
	})

	t.Run("update existing key", func(t *testing.T) {
		cache := NewLRUCache(2)
		if cache == nil {
			t.Fatal("NewLRUCache returned nil")
		}

		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Put(1, 10) // Update key 1

		if got := cache.Get(1); got != 10 {
			t.Errorf("Get(1) = %d, want 10", got)
		}

		// Key 2 should still be there
		if got := cache.Get(2); got != 2 {
			t.Errorf("Get(2) = %d, want 2", got)
		}
	})

	t.Run("get refreshes recency", func(t *testing.T) {
		cache := NewLRUCache(2)
		if cache == nil {
			t.Fatal("NewLRUCache returned nil")
		}

		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Get(1) // Access key 1, making key 2 the LRU

		cache.Put(3, 3) // Should evict key 2

		if got := cache.Get(1); got != 1 {
			t.Errorf("Get(1) = %d, want 1", got)
		}

		if got := cache.Get(2); got != -1 {
			t.Errorf("Get(2) = %d, want -1 (evicted)", got)
		}
	})

	t.Run("capacity of 1", func(t *testing.T) {
		cache := NewLRUCache(1)
		if cache == nil {
			t.Fatal("NewLRUCache returned nil")
		}

		cache.Put(1, 1)
		cache.Put(2, 2)

		if got := cache.Get(1); got != -1 {
			t.Errorf("Get(1) = %d, want -1 (evicted)", got)
		}

		if got := cache.Get(2); got != 2 {
			t.Errorf("Get(2) = %d, want 2", got)
		}
	})
}
