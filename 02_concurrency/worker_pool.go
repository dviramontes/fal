package concurrency

import "context"

// Exercise 1: Worker Pool Pattern
// Implement a worker pool that processes stock analysis jobs concurrently.
// This is fundamental for distributed systems at scale.
//
// Requirements:
// - Create a pool of N workers
// - Jobs are submitted to an input channel
// - Results are collected from an output channel
// - Workers should stop gracefully when context is cancelled
//
// Example:
//   pool := NewWorkerPool(4)
//   results := pool.Process(ctx, jobs, analyzeStock)

type Job struct {
	ID     int
	Ticker string
	Data   []float64
}

type Result struct {
	JobID  int
	Ticker string
	Output float64
	Err    error
}

type WorkerPool struct {
	// TODO: Add fields
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	// TODO: Implement this
	return nil
}

// Process takes a channel of jobs and a processing function,
// returns a channel of results
func (wp *WorkerPool) Process(
	ctx context.Context,
	jobs <-chan Job,
	processor func(Job) Result,
) <-chan Result {
	// TODO: Implement this
	return nil
}

// Exercise 2: Rate Limiter
// Implement a token bucket rate limiter for API calls.
// Essential for respecting exchange API limits.
//
// Requirements:
// - Allow burst of up to 'burst' requests
// - Refill at 'rate' tokens per second
// - Wait() blocks until a token is available or context is cancelled
//
// Example:
//   limiter := NewRateLimiter(10, 100) // 10 per second, burst of 100
//   limiter.Wait(ctx) // blocks if no tokens available

type RateLimiter struct {
	// TODO: Add fields
}

func NewRateLimiter(rate float64, burst int) *RateLimiter {
	// TODO: Implement this
	return nil
}

func (rl *RateLimiter) Wait(ctx context.Context) error {
	// TODO: Implement this
	return nil
}

func (rl *RateLimiter) Allow() bool {
	// TODO: Implement this - non-blocking version
	return false
}
