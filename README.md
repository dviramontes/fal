# Falcon

Interview practice in Go, Odin, and Elixir, covering data structures, concurrency, data processing, and financial algorithms. The goal is hands-on language practice across a broad range of problems.

## Setup

Install the tools for the language you want to practice:

| Language | Requirements | Setup |
| --- | --- | --- |
| Go | Go 1.23+ | From `go/`, run `go mod download` and `go install gotest.tools/gotestsum@latest` |
| Odin | [Odin](https://odin-lang.org/docs/install/) | No dependency setup required |
| Elixir | See [Elixir setup](elixir/README.md#setup) | See the language README |

Install [just](https://github.com/casey/just) to use the command recipes below.

## Project Structure

```
falcon/
├── go/
│   ├── data_structures/    # Classic data structure problems
│   ├── concurrency/        # Go concurrency patterns
│   ├── data_processing/    # Stream and batch processing
│   ├── financial/          # Financial domain algorithms
│   ├── justfile            # Go commands
│   ├── go.mod
│   └── go.sum
├── odin/
│   ├── data_structures/    # Grid and data structure problems
│   └── justfile            # Odin commands
├── elixir/
│   ├── lib/                # Elixir modules
│   ├── test/               # ExUnit tests
│   ├── justfile            # Elixir commands
│   └── mix.exs             # Mix project (no external dependencies)
└── README.md
```

## Running Tests

Each language has its own `justfile`; there is no root justfile. Run `just test` from `go/`, `odin/`, or `elixir/`. Direct language commands also run from their respective directories.

From the repository root, select a justfile explicitly:

```bash
just --justfile go/justfile test
just --justfile odin/justfile test
just --justfile elixir/justfile test
```

In any language directory, `just --list` lists commands and `just exercises` lists exercises. Tests for unsolved exercises are expected to fail.

### Go Commands

Go tests use [testify](https://github.com/stretchr/testify) assertions and [gotestsum](https://github.com/gotestyourself/gotestsum) for test output and watch mode.

Run these from `go/`:

```bash
just test-dots         # Compact test output
just test-ds           # Data structures
just test-concurrency
just test-data         # Data processing
just test-financial
just test-one TestLRUCache
just test-race
just watch             # Rerun all tests on changes
just watch-ds          # Watch data structure tests
just test-coverage     # Generate coverage report
just bench
just fmt
just lint              # Requires golangci-lint
```

### Odin Commands

Run `just test-ds` from `odin/` to test data structure exercises with Odin's built-in test runner.

### Elixir Commands

See the [Elixir command reference](elixir/README.md#commands) for ExUnit tests, formatting, compilation checks, and interactive sessions.

## Practice Workflow

1. Pick an exercise from `just exercises` in the language directory.
2. Read its problem description and hints.
3. Run its tests to see the expected behavior.
4. Implement the solution and rerun the tests as you work. For Go, watch mode automates reruns.

## Go Exercises

File names below are relative to their topic directory under `go/`. Source files contain the problem statements; adjacent `*_test.go` files define expected behavior.

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

## Odin Exercises

[Minutes Until Apocalypse](odin/data_structures/minutes_until_apocalypse.odin) is a grid infection exercise. Implement `minutes_until_apocalypse(grid: [][]int) -> int` to satisfy the adjacent tests, which cover simultaneous infection, unreachable people, empty grids, and row/column boundaries.

## Elixir Exercises

See the [Elixir exercise catalog](elixir/README.md#exercises).

## License

MIT
