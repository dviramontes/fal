# BAI - Go Exercises

# Default recipe - show available commands
default:
    @just --list

# Run all tests
test:
    gotestsum --format testname ./...

# Run all tests (compact dots format)
test-dots:
    gotestsum --format dots ./...

# Run tests with coverage
test-coverage:
    go test ./... -coverprofile=coverage.out
    go tool cover -html=coverage.out -o coverage.html
    @echo "Coverage report: coverage.html"

# Run a specific exercise module
test-ds:
    gotestsum --format testname ./01_data_structures/...

test-concurrency:
    gotestsum --format testname ./02_concurrency/...

test-data:
    gotestsum --format testname ./03_data_processing/...

test-financial:
    gotestsum --format testname ./04_financial/...

# Run a single test by name (e.g., just test-one TestLRUCache)
test-one NAME:
    gotestsum --format testname -- -run {{NAME}} ./...

# Run tests in short mode (skip long-running tests)
test-short:
    gotestsum --format testname -- -short ./...

# Run benchmarks
bench:
    go test ./... -bench=. -benchmem

# Check for race conditions
test-race:
    gotestsum --format testname -- -race ./...

# Format all code
fmt:
    go fmt ./...

# Lint the code (requires golangci-lint)
lint:
    golangci-lint run

# Tidy dependencies
tidy:
    go mod tidy

# Build to check for compilation errors
build:
    go build ./...

# Clean generated files
clean:
    rm -f coverage.out coverage.html

# Watch tests (uses gotestsum's built-in watch mode)
watch:
    gotestsum --watch --format testname ./...

# Practice workflow: run single module tests in watch mode
watch-ds:
    gotestsum --watch --format testname ./01_data_structures/...

watch-concurrency:
    gotestsum --watch --format testname ./02_concurrency/...

watch-data:
    gotestsum --watch --format testname ./03_data_processing/...

watch-financial:
    gotestsum --watch --format testname ./04_financial/...

# Show exercise summary
exercises:
    @echo ""
    @echo "=== BAI PRACTICE ==="
    @echo ""
    @echo "01_data_structures/"
    @echo "  - Sliding Window Maximum (sliding_window.go)"
    @echo "  - Moving Average Stream (sliding_window.go)"
    @echo "  - LRU Cache (lru_cache.go)"
    @echo "  - Top K Frequent Elements (heap_topk.go)"
    @echo "  - Merge K Sorted Streams (heap_topk.go)"
    @echo "  - Find Median from Stream (heap_topk.go)"
    @echo "  - Count Greater Than Previous Average (average_counter.go)"
    @echo ""
    @echo "02_concurrency/"
    @echo "  - Worker Pool Pattern (worker_pool.go)"
    @echo "  - Rate Limiter (worker_pool.go)"
    @echo "  - Fan-Out Fan-In (fanout_fanin.go)"
    @echo "  - Pipeline Pattern (fanout_fanin.go)"
    @echo "  - Semaphore (fanout_fanin.go)"
    @echo "  - Pub/Sub Broker (pubsub.go)"
    @echo "  - Circuit Breaker (pubsub.go)"
    @echo ""
    @echo "03_data_processing/"
    @echo "  - Time-Windowed Aggregation (stream_processor.go)"
    @echo "  - Rolling Statistics (stream_processor.go)"
    @echo "  - Data Deduplication (stream_processor.go)"
    @echo "  - Batch Processor (batch_processor.go)"
    @echo "  - Merge Sorted Iterators (batch_processor.go)"
    @echo "  - Event Sourcing Store (batch_processor.go)"
    @echo ""
    @echo "04_financial/"
    @echo "  - Portfolio Rebalancing (portfolio.go)"
    @echo "  - Order Book (portfolio.go)"
    @echo "  - VWAP Calculator (portfolio.go)"
    @echo "  - Sharpe Ratio (risk.go)"
    @echo "  - Maximum Drawdown (risk.go)"
    @echo "  - Kelly Criterion (risk.go)"
    @echo "  - Correlation Matrix (risk.go)"
    @echo "  - Technical Indicators: SMA, EMA, RSI, Bollinger (timeseries.go)"
    @echo "  - Return Calculations (timeseries.go)"
    @echo ""
    @echo "Run 'just test' to see which exercises pass/fail"
    @echo ""
