# Boosted.ai Interview Practice - Go Exercises

# Default recipe - show available commands
default:
    @just --list

# Run all tests
test:
    go test ./... -v

# Run tests with coverage
test-coverage:
    go test ./... -coverprofile=coverage.out
    go tool cover -html=coverage.out -o coverage.html
    @echo "Coverage report: coverage.html"

# Run a specific exercise module
test-ds:
    @echo "=== Data Structures & Algorithms ==="
    go test ./01_data_structures/... -v

test-concurrency:
    @echo "=== Concurrency Patterns ==="
    go test ./02_concurrency/... -v

test-data:
    @echo "=== Data Processing ==="
    go test ./03_data_processing/... -v

test-financial:
    @echo "=== Financial Domain ==="
    go test ./04_financial/... -v

# Run a single test by name (e.g., just test-one TestLRUCache)
test-one NAME:
    go test ./... -v -run {{NAME}}

# Run tests in short mode (skip long-running tests)
test-short:
    go test ./... -v -short

# Run benchmarks
bench:
    go test ./... -bench=. -benchmem

# Check for race conditions
test-race:
    go test ./... -race

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

# Watch tests (requires entr: brew install entr)
watch:
    find . -name '*.go' | entr -c go test ./... -v

# Practice workflow: run single module tests in watch mode
watch-ds:
    find ./01_data_structures -name '*.go' | entr -c go test ./01_data_structures/... -v

watch-concurrency:
    find ./02_concurrency -name '*.go' | entr -c go test ./02_concurrency/... -v

watch-data:
    find ./03_data_processing -name '*.go' | entr -c go test ./03_data_processing/... -v

watch-financial:
    find ./04_financial -name '*.go' | entr -c go test ./04_financial/... -v

# Show exercise summary
exercises:
    @echo ""
    @echo "=== BOOSTED.AI INTERVIEW PRACTICE ==="
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
