package concurrency

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	t.Run("processes all jobs", func(t *testing.T) {
		pool := NewWorkerPool(4)
		if pool == nil {
			t.Fatal("NewWorkerPool returned nil")
		}

		ctx := context.Background()
		jobs := make(chan Job, 10)

		// Send jobs
		for i := 0; i < 10; i++ {
			jobs <- Job{ID: i, Ticker: "AAPL", Data: []float64{float64(i)}}
		}
		close(jobs)

		processor := func(j Job) Result {
			return Result{JobID: j.ID, Ticker: j.Ticker, Output: j.Data[0] * 2}
		}

		results := pool.Process(ctx, jobs, processor)

		var count int
		for result := range results {
			if result.Output != float64(result.JobID)*2 {
				t.Errorf("Job %d: got output %f, want %f",
					result.JobID, result.Output, float64(result.JobID)*2)
			}
			count++
		}

		if count != 10 {
			t.Errorf("Processed %d jobs, want 10", count)
		}
	})

	t.Run("respects context cancellation", func(t *testing.T) {
		pool := NewWorkerPool(2)
		if pool == nil {
			t.Fatal("NewWorkerPool returned nil")
		}

		ctx, cancel := context.WithCancel(context.Background())
		jobs := make(chan Job, 100)

		// Send many jobs
		for i := 0; i < 100; i++ {
			jobs <- Job{ID: i, Ticker: "AAPL"}
		}
		close(jobs)

		processor := func(j Job) Result {
			time.Sleep(10 * time.Millisecond)
			return Result{JobID: j.ID}
		}

		results := pool.Process(ctx, jobs, processor)

		// Cancel after receiving some results
		go func() {
			time.Sleep(50 * time.Millisecond)
			cancel()
		}()

		var count int
		for range results {
			count++
		}

		// Should have processed fewer than all jobs
		if count >= 100 {
			t.Error("Context cancellation did not stop processing")
		}
	})
}

func TestRateLimiter(t *testing.T) {
	t.Run("allows burst", func(t *testing.T) {
		limiter := NewRateLimiter(10, 5) // 10/sec, burst of 5
		if limiter == nil {
			t.Fatal("NewRateLimiter returned nil")
		}

		// Should allow burst of 5 immediately
		for i := 0; i < 5; i++ {
			if !limiter.Allow() {
				t.Errorf("Allow() = false on request %d, want true (within burst)", i+1)
			}
		}
	})

	t.Run("blocks after burst", func(t *testing.T) {
		limiter := NewRateLimiter(10, 3)
		if limiter == nil {
			t.Fatal("NewRateLimiter returned nil")
		}

		// Exhaust burst
		for i := 0; i < 3; i++ {
			limiter.Allow()
		}

		// Should be denied
		if limiter.Allow() {
			t.Error("Allow() = true after burst exhausted, want false")
		}
	})

	t.Run("Wait respects context", func(t *testing.T) {
		limiter := NewRateLimiter(1, 1) // 1/sec
		if limiter == nil {
			t.Fatal("NewRateLimiter returned nil")
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel()

		// Use the one token
		limiter.Allow()

		// Should fail due to context timeout
		err := limiter.Wait(ctx)
		if err == nil {
			t.Error("Wait() returned nil error, want context deadline exceeded")
		}
	})
}

func TestSemaphore(t *testing.T) {
	t.Run("limits concurrency", func(t *testing.T) {
		sem := NewSemaphore(3)
		if sem == nil {
			t.Fatal("NewSemaphore returned nil")
		}

		var (
			current int
			maxSeen int
			mu      sync.Mutex
			wg      sync.WaitGroup
		)

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				ctx := context.Background()
				if err := sem.Acquire(ctx); err != nil {
					return
				}

				mu.Lock()
				current++
				if current > maxSeen {
					maxSeen = current
				}
				mu.Unlock()

				time.Sleep(10 * time.Millisecond)

				mu.Lock()
				current--
				mu.Unlock()

				sem.Release()
			}()
		}

		wg.Wait()

		if maxSeen > 3 {
			t.Errorf("Max concurrent = %d, want <= 3", maxSeen)
		}
	})
}
