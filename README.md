# Falcon

A collection of programming exercises organized by language, focused on data structures, concurrency patterns, and financial domain problems. Go exercises live in `go/`; `odin/` is reserved for future exercises and is currently empty. Each Go exercise includes comprehensive tests - implement the functions to make the tests pass.

## Prerequisites

- Go 1.23+
- [just](https://github.com/casey/just) command runner (recommended)
- [gotestsum](https://github.com/gotestyourself/gotestsum) for enhanced test output

## Setup

```bash
# Clone the repository
git clone <repo-url>
cd falcon

# Install dependencies
just tidy

# Install gotestsum (for better test output)
go install gotest.tools/gotestsum@latest
```

## Project Structure

```
falcon/
├── go/
│   ├── data_structures/    # Classic data structure problems
│   ├── concurrency/        # Go concurrency patterns
│   ├── data_processing/    # Stream and batch processing
│   ├── financial/          # Financial domain algorithms
│   ├── go.mod
│   └── go.sum
├── odin/                   # Empty; reserved for future exercises
├── justfile                # Task runner commands
└── README.md
```

Git does not track empty directories, so `odin/` will appear in fresh clones once it contains files.

## Exercises

### Data Structures & Algorithms

| Exercise | File | Concepts |
|----------|------|----------|
| Sliding Window Maximum | `sliding_window.go` | Deque, O(n) window operations |
| Moving Average Stream | `sliding_window.go` | Circular buffer, streaming data |
| LRU Cache | `lru_cache.go` | Hash map + doubly linked list |
| Top K Frequent Elements | `heap_topk.go` | Min-heap, frequency counting |
| Merge K Sorted Streams | `heap_topk.go` | Min-heap, k-way merge |
| Find Median from Stream | `heap_topk.go` | Two heaps (max + min) |
| Count Greater Than Average | `average_counter.go` | Running statistics |

### Concurrency Patterns

| Exercise | File | Concepts |
|----------|------|----------|
| Worker Pool | `worker_pool.go` | Goroutines, channels, fan-out |
| Rate Limiter | `worker_pool.go` | Token bucket algorithm |
| Semaphore | `fanout_fanin.go` | Bounded concurrency |
| Fan-Out Fan-In | `fanout_fanin.go` | Pipeline patterns |
| Pub/Sub Broker | `pubsub.go` | Event-driven architecture |
| Circuit Breaker | `pubsub.go` | Fault tolerance |

### Data Processing

| Exercise | File | Concepts |
|----------|------|----------|
| Window Aggregator | `stream_processor.go` | Time-based windowing, OHLCV |
| Rolling Statistics | `stream_processor.go` | Mean, std dev, min/max |
| Deduplicator | `stream_processor.go` | Bloom filter concepts |
| Batch Processor | `batch_processor.go` | Batching, context cancellation |
| Event Store | `batch_processor.go` | Event sourcing, snapshots |

### Financial Domain

| Exercise | File | Concepts |
|----------|------|----------|
| Portfolio Rebalancing | `portfolio.go` | Allocation optimization |
| Order Book | `portfolio.go` | Price-time priority matching |
| VWAP Calculator | `portfolio.go` | Volume-weighted average |
| Sharpe Ratio | `risk.go` | Risk-adjusted returns |
| Maximum Drawdown | `risk.go` | Peak-to-trough analysis |
| Kelly Criterion | `risk.go` | Optimal position sizing |
| Correlation Matrix | `risk.go` | Pairwise correlations |
| Technical Indicators | `timeseries.go` | SMA, EMA, RSI, Bollinger Bands |
| Return Calculations | `timeseries.go` | Simple, log, cumulative returns |

## Running Tests

This project uses [gotestsum](https://github.com/gotestyourself/gotestsum) for improved test output with better formatting, colors, and failure summaries.

Run `just` commands from the repository root; recipes use `go/` as their working directory. For direct Go commands, first run `cd go`. Exercise file paths below are relative to their topic directory under `go/`.

### Quick Reference

```bash
# Run all tests
just test
just test-go          # Explicit Go alias

# Run tests (compact dots format)
just test-dots

# Run specific module
just test-ds           # Data structures
just test-concurrency  # Concurrency
just test-data         # Data processing
just test-financial    # Financial

# Run a single test by name
just test-one TestLRUCache

# Run with race detector
just test-race

# Run in watch mode (re-runs on file changes)
just watch
just watch-ds          # Watch specific module
```

### All Commands

```bash
just --list            # Show all available commands
just test-coverage     # Generate coverage report
just bench             # Run benchmarks
just fmt               # Format code
just lint              # Lint code (requires golangci-lint)
just exercises         # Show exercise summary
```

## Workflow

1. Pick an exercise from `just exercises`
2. Read the problem description and hints in the source file
3. Start watch mode: `just watch-ds` (or appropriate module)
4. Implement the solution
5. Tests will re-run automatically on save
6. Repeat until all tests pass

## Testing Libraries

This project uses:

- **[testify](https://github.com/stretchr/testify)** - Better assertions with clear failure messages
- **[gotestsum](https://github.com/gotestyourself/gotestsum)** - Enhanced test runner with:
  - Colored pass/fail output
  - Test name formatting
  - Summary statistics
  - Built-in watch mode

## License

MIT
