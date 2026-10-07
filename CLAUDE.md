# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go interview practice exercises covering data structures, concurrency patterns, and financial domain problems. Each exercise has tests - implement the functions to make them pass.

Exercises are organized by language: `go/` contains the Go module and exercises; `odin/` is currently empty and reserved for future exercises.

## Commands

```bash
# Run all tests
just test
just test-go           # Explicit Go alias

# Run specific module tests
just test-ds           # data_structures
just test-concurrency  # concurrency
just test-data         # data_processing
just test-financial    # financial

# Run a single test by name
just test-one TestLRUCache

# Watch mode (re-runs on file changes)
just watch
just watch-ds          # Watch specific module

# Other
just test-race         # Race condition detection
just fmt               # Format code
just lint              # Lint (requires golangci-lint)
just bench             # Run benchmarks
```

## Architecture

Module name: `falcon` (Go 1.23+), rooted at `go/`. Run `just` commands from the repository root; recipes execute in `go/`. Run direct Go commands from `go/`.

Each topic directory under `go/` is a separate package:
- `datastructures` - Classic DS problems (sliding window, LRU cache, heaps)
- `concurrency` - Go patterns (worker pool, pub/sub, circuit breaker)
- `dataprocessing` - Stream/batch processing, event sourcing
- `financial` - Portfolio, risk metrics, technical indicators

Tests use [testify](https://github.com/stretchr/testify) for assertions. Test runner is [gotestsum](https://github.com/gotestyourself/gotestsum).

## Exercise Pattern

Each `.go` file contains:
1. Problem description in comments
2. Constraints and examples
3. Implementation hints
4. Stub functions to implement

## Goals

The primary goal of this project is to get as much practice as I can with the language
before my interview so I want to work on as many types of problems as possible

Corresponding `_test.go` files have comprehensive test cases.

## Assistance Pattern

- Provide codee review when ask but don't implement solutions
- Ignore minor issues or style comments
- Focus on implementation details but don't over do it
