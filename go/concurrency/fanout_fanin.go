package concurrency

import "context"

// Exercise 3: Fan-Out Fan-In Pattern
// Process market data by fanning out to multiple analyzers,
// then fan in the results. Common for parallel data processing pipelines.
//
// Requirements:
// - Fan out: distribute work to N goroutines
// - Fan in: merge results from all goroutines into single channel
// - Handle context cancellation gracefully

type MarketData struct {
	Ticker    string
	Price     float64
	Volume    int64
	Timestamp int64
}

type AnalysisResult struct {
	Ticker     string
	Signal     string // "BUY", "SELL", "HOLD"
	Confidence float64
}

// FanOut distributes data from input channel to N output channels
func FanOut(ctx context.Context, input <-chan MarketData, n int) []<-chan MarketData {
	// TODO: Implement this
	return nil
}

// FanIn merges multiple input channels into a single output channel
func FanIn(ctx context.Context, inputs ...<-chan AnalysisResult) <-chan AnalysisResult {
	// TODO: Implement this
	return nil
}

// Exercise 4: Pipeline Pattern
// Build a data processing pipeline with stages:
// 1. Fetch raw data
// 2. Clean/normalize data
// 3. Calculate indicators
// 4. Generate signals
//
// Each stage runs concurrently, connected by channels.

type PipelineStage[In, Out any] func(ctx context.Context, input <-chan In) <-chan Out

// Pipeline chains multiple stages together
func Pipeline[T any](ctx context.Context, source <-chan T, stages ...PipelineStage[T, T]) <-chan T {
	// TODO: Implement this
	// Note: This simplified version assumes same type throughout
	// A real implementation might use interface{} or separate functions
	return nil
}

// Exercise 5: Semaphore
// Implement a counting semaphore to limit concurrent database connections.
//
// Example:
//   sem := NewSemaphore(10) // max 10 concurrent connections
//   sem.Acquire(ctx)
//   defer sem.Release()
//   // do database work

type Semaphore struct {
	// TODO: Add fields
}

func NewSemaphore(max int) *Semaphore {
	// TODO: Implement this
	return nil
}

func (s *Semaphore) Acquire(ctx context.Context) error {
	// TODO: Implement this
	return nil
}

func (s *Semaphore) Release() {
	// TODO: Implement this
}
