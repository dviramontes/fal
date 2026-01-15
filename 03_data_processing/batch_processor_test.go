package dataprocessing

import (
	"context"
	"testing"
)

func TestBatchProcessor(t *testing.T) {
	t.Run("processes all items", func(t *testing.T) {
		bp := NewBatchProcessor[int, int](BatchConfig{
			BatchSize:   3,
			Concurrency: 2,
		})
		if bp == nil {
			t.Fatal("NewBatchProcessor returned nil")
		}

		items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		results, err := bp.Process(context.Background(), items, func(batch []int) ([]int, error) {
			out := make([]int, len(batch))
			for i, v := range batch {
				out[i] = v * 2
			}
			return out, nil
		})

		if err != nil {
			t.Fatalf("Process returned error: %v", err)
		}

		if len(results) != len(items) {
			t.Errorf("Got %d results, want %d", len(results), len(items))
		}

		// Check all results are doubled (order may vary)
		sum := 0
		for _, r := range results {
			sum += r
		}
		expectedSum := 2 * (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10)
		if sum != expectedSum {
			t.Errorf("Sum of results = %d, want %d", sum, expectedSum)
		}
	})

	t.Run("respects context cancellation", func(t *testing.T) {
		bp := NewBatchProcessor[int, int](BatchConfig{
			BatchSize:   2,
			Concurrency: 1,
		})
		if bp == nil {
			t.Fatal("NewBatchProcessor returned nil")
		}

		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		items := []int{1, 2, 3, 4, 5}
		_, err := bp.Process(ctx, items, func(batch []int) ([]int, error) {
			return batch, nil
		})

		if err == nil {
			t.Error("Expected error due to cancelled context")
		}
	})
}

func TestEventStore(t *testing.T) {
	t.Run("append and retrieve events", func(t *testing.T) {
		store := NewEventStore()
		if store == nil {
			t.Fatal("NewEventStore returned nil")
		}

		events := []Event{
			{AggregateID: "order-1", Type: "OrderCreated", Version: 1, Data: map[string]interface{}{"amount": 100}},
			{AggregateID: "order-1", Type: "OrderPaid", Version: 2, Data: map[string]interface{}{"paidAt": 12345}},
			{AggregateID: "order-2", Type: "OrderCreated", Version: 1, Data: map[string]interface{}{"amount": 200}},
		}

		for _, e := range events {
			if err := store.Append(e); err != nil {
				t.Fatalf("Append failed: %v", err)
			}
		}

		retrieved, err := store.GetEvents("order-1", 0)
		if err != nil {
			t.Fatalf("GetEvents failed: %v", err)
		}

		if len(retrieved) != 2 {
			t.Errorf("Got %d events, want 2", len(retrieved))
		}
	})

	t.Run("retrieve events from version", func(t *testing.T) {
		store := NewEventStore()
		if store == nil {
			t.Fatal("NewEventStore returned nil")
		}

		for i := 1; i <= 5; i++ {
			store.Append(Event{AggregateID: "agg-1", Version: int64(i)})
		}

		retrieved, _ := store.GetEvents("agg-1", 3)
		if len(retrieved) != 3 { // versions 3, 4, 5
			t.Errorf("Got %d events from version 3, want 3", len(retrieved))
		}
	})

	t.Run("snapshots", func(t *testing.T) {
		store := NewEventStore()
		if store == nil {
			t.Fatal("NewEventStore returned nil")
		}

		snap := Snapshot{
			AggregateID: "agg-1",
			Version:     10,
			State:       map[string]interface{}{"balance": 500},
		}

		if err := store.SaveSnapshot(snap); err != nil {
			t.Fatalf("SaveSnapshot failed: %v", err)
		}

		retrieved, err := store.GetSnapshot("agg-1")
		if err != nil {
			t.Fatalf("GetSnapshot failed: %v", err)
		}

		if retrieved.Version != 10 {
			t.Errorf("Snapshot version = %d, want 10", retrieved.Version)
		}
	})
}
